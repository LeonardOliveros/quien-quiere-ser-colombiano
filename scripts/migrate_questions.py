#!/usr/bin/env python3
"""One-time migration: data/questions.json -> data/taxonomy.json + data/questions/*.json.

- Consolidates the 156 free-text subcategories into a curated taxonomy (~44 codes).
- Drops questions with duplicate normalized text (keeps the copy with an explanation).
- Assigns each question a stable key (CUL-0001, GEO-0001, ...) used by the seeder
  to upsert questions idempotently. Keys must never be reassigned after this run.

Kept in the repo for traceability; it is not needed at runtime.
"""

import json
import unicodedata
from collections import Counter
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
SOURCE = ROOT / "data" / "questions.json"
TAXONOMY_OUT = ROOT / "data" / "taxonomy.json"
QUESTIONS_DIR = ROOT / "data" / "questions"

# Canonical taxonomy: category code -> (display name, {subcategory code: display name})
TAXONOMY = {
    "CULTURA": ("Cultura", {
        "GENERAL": "General",
        "SIMBOLOS_PATRIOS": "Símbolos Patrios",
        "ARTE_Y_ARTISTAS": "Arte y Artistas",
        "ARQUITECTURA_Y_TEATROS": "Arquitectura y Teatros",
        "LITERATURA": "Literatura",
        "MUSICA_Y_DANZA": "Música y Danza",
        "CINE_Y_TELEVISION": "Cine y Televisión",
        "DEPORTES": "Deportes",
        "FESTIVALES_Y_FERIAS": "Festivales y Ferias",
        "GASTRONOMIA": "Gastronomía",
        "TURISMO": "Turismo",
        "MONEDA_Y_BILLETES": "Moneda y Billetes",
        "RELIGION_Y_TRADICIONES": "Religión y Tradiciones",
        "ETNIAS_Y_DEMOGRAFIA": "Etnias y Demografía",
    }),
    "GEOGRAFIA": ("Geografía", {
        "GENERAL": "General",
        "FRONTERAS_Y_LIMITES": "Fronteras y Límites",
        "DIVISION_POLITICA": "División Política",
        "REGIONES_NATURALES": "Regiones Naturales",
        "RELIEVE": "Relieve",
        "HIDROGRAFIA": "Hidrografía",
        "CLIMA_Y_PISOS_TERMICOS": "Clima y Pisos Térmicos",
        "FLORA_Y_FAUNA": "Flora y Fauna",
        "PARQUES_NACIONALES": "Parques Nacionales",
        "ECONOMIA_Y_RECURSOS": "Economía y Recursos",
        "POBLACION_Y_ETNIAS": "Población y Etnias",
    }),
    "HISTORIA": ("Historia", {
        "GENERAL": "General",
        "PRECOLOMBINO": "Período Precolombino",
        "CONQUISTA": "La Conquista",
        "COLONIA": "La Colonia",
        "INDEPENDENCIA": "La Independencia",
        "GRAN_COLOMBIA": "La Gran Colombia",
        "SIGLO_XIX": "Siglo XIX",
        "SIGLO_XX": "Siglo XX",
        "CONFLICTO_Y_PAZ": "Conflicto Armado y Paz",
    }),
    "CONSTITUCION": ("Constitución", {
        "GENERAL": "General",
        "PRINCIPIOS_FUNDAMENTALES": "Principios Fundamentales",
        "DERECHOS_Y_LIBERTADES": "Derechos y Libertades",
        "MECANISMOS_DE_PROTECCION": "Mecanismos de Protección",
        "PARTICIPACION_Y_ELECCIONES": "Participación y Elecciones",
        "RAMAS_DEL_PODER_PUBLICO": "Ramas del Poder Público",
        "ORGANISMOS_DE_CONTROL": "Organismos de Control",
        "ORGANIZACION_TERRITORIAL": "Organización Territorial",
        "ESTADOS_DE_EXCEPCION": "Estados de Excepción",
        "REFORMAS_CONSTITUCIONALES": "Reformas Constitucionales",
    }),
}

# Old free-text subcategory -> canonical code, per category. Every value present
# in the source file must appear here; the script fails loudly otherwise.
MAPPING = {
    "CULTURA": {
        "": "GENERAL",
        "HISTORIA": "GENERAL",
        "INFRAESTRUCTURA": "GENERAL",
        "INSTITUCIONES": "GENERAL",
        "LEGISLACIÓN": "GENERAL",
        "LITORAL": "GENERAL",
        "NATURALEZA": "GENERAL",
        "REGIONES": "GENERAL",
        "RURAL": "GENERAL",
        "AVE NACIONAL": "SIMBOLOS_PATRIOS",
        "BANDERA NACIONAL": "SIMBOLOS_PATRIOS",
        "ESCUDO NACIONAL": "SIMBOLOS_PATRIOS",
        "FLOR NACIONAL": "SIMBOLOS_PATRIOS",
        "HIMNO NACIONAL": "SIMBOLOS_PATRIOS",
        "NOMBRE OFICIAL": "SIMBOLOS_PATRIOS",
        "SÍMBOLO CULTURAL": "SIMBOLOS_PATRIOS",
        "SÍMBOLOS PATRIOS": "SIMBOLOS_PATRIOS",
        "ÁRBOL NACIONAL": "SIMBOLOS_PATRIOS",
        "ARTE Y ARTISTAS ESPECÍFICOS": "ARTE_Y_ARTISTAS",
        "ARTES PLÁSTICAS": "ARTE_Y_ARTISTAS",
        "ARTES Y CULTURA": "ARTE_Y_ARTISTAS",
        "ARQUITECTURA": "ARQUITECTURA_Y_TEATROS",
        "ARQUITECTURA Y TEATROS": "ARQUITECTURA_Y_TEATROS",
        "TEATRO": "ARQUITECTURA_Y_TEATROS",
        "LITERATURA": "LITERATURA",
        "LITERATURA Y ESCRITORES": "LITERATURA",
        "MÚSICA Y DANZA": "MUSICA_Y_DANZA",
        "MÚSICA Y DANZA REGIONAL": "MUSICA_Y_DANZA",
        "CINE": "CINE_Y_TELEVISION",
        "RADIO Y TELEVISIÓN": "CINE_Y_TELEVISION",
        "DEPORTE": "DEPORTES",
        "DEPORTES": "DEPORTES",
        "DEPORTES ESPECÍFICOS": "DEPORTES",
        "FESTIVALES": "FESTIVALES_Y_FERIAS",
        "FESTIVALES Y CELEBRACIONES": "FESTIVALES_Y_FERIAS",
        "FESTIVIDADES Y FERIAS REGIONALES": "FESTIVALES_Y_FERIAS",
        "GASTRONOMÍA": "GASTRONOMIA",
        "GASTRONOMÍA REGIONAL": "GASTRONOMIA",
        "GASTRONOMÍA REGIONAL DETALLADA": "GASTRONOMIA",
        "MÚSICA Y GASTRONOMÍA": "GASTRONOMIA",
        "TURISMO": "TURISMO",
        "TURISMO Y SITIOS ESPECÍFICOS": "TURISMO",
        "MONEDA": "MONEDA_Y_BILLETES",
        "MONEDA Y BILLETES": "MONEDA_Y_BILLETES",
        "RELIGIÓN": "RELIGION_Y_TRADICIONES",
        "RELIGIÓN Y TRADICIONES": "RELIGION_Y_TRADICIONES",
        "ETNIAS Y DEMOGRAFÍA": "ETNIAS_Y_DEMOGRAFIA",
    },
    "GEOGRAFIA": {
        "": "GENERAL",
        "GENERAL": "GENERAL",
        "NACIONAL": "GENERAL",
        "FRONTERAS": "FRONTERAS_Y_LIMITES",
        "FRONTERAS Y LÍMITES": "FRONTERAS_Y_LIMITES",
        "UBICACIÓN": "FRONTERAS_Y_LIMITES",
        "UBICACIÓN Y LÍMITES": "FRONTERAS_Y_LIMITES",
        "BOGOTÁ": "DIVISION_POLITICA",
        "CALI": "DIVISION_POLITICA",
        "MEDELLÍN": "DIVISION_POLITICA",
        "CAPITALES": "DIVISION_POLITICA",
        "DEPARTAMENTOS": "DIVISION_POLITICA",
        "DEPARTAMENTOS Y CAPITALES": "DIVISION_POLITICA",
        "DIVISION POLITICA": "DIVISION_POLITICA",
        "DIVISIÓN POLÍTICA": "DIVISION_POLITICA",
        "REGIONES NATURALES": "REGIONES_NATURALES",
        "REGIÓN AMAZÓNICA": "REGIONES_NATURALES",
        "REGIÓN AMAZÓNICA - DETALLES": "REGIONES_NATURALES",
        "REGIÓN ANDINA - DETALLES": "REGIONES_NATURALES",
        "REGIÓN CARIBE - CARACTERÍSTICAS ESPECÍFICAS": "REGIONES_NATURALES",
        "REGIÓN INSULAR": "REGIONES_NATURALES",
        "REGIÓN ORINOQUÍA - CARACTERÍSTICAS": "REGIONES_NATURALES",
        "REGIÓN PACÍFICA - CARACTERÍSTICAS": "REGIONES_NATURALES",
        "TERRITORIO INSULAR": "REGIONES_NATURALES",
        "RELIEVE": "RELIEVE",
        "SISTEMA OROGRÁFICO": "RELIEVE",
        "LLANURAS": "RELIEVE",
        "VALLES INTERANDINOS": "RELIEVE",
        "HIDROGRAFÍA": "HIDROGRAFIA",
        "CLIMA Y BIODIVERSIDAD": "CLIMA_Y_PISOS_TERMICOS",
        "PISOS TÉRMICOS": "CLIMA_Y_PISOS_TERMICOS",
        "PISOS TÉRMICOS Y CULTIVOS": "CLIMA_Y_PISOS_TERMICOS",
        "ECOSISTEMAS": "FLORA_Y_FAUNA",
        "FAUNA": "FLORA_Y_FAUNA",
        "FLORA": "FLORA_Y_FAUNA",
        "MAMÍFEROS": "FLORA_Y_FAUNA",
        "REPTILES": "FLORA_Y_FAUNA",
        "PARQUES NACIONALES": "PARQUES_NACIONALES",
        "ACTIVIDADES PRODUCTIVAS": "ECONOMIA_Y_RECURSOS",
        "AGRICULTURA": "ECONOMIA_Y_RECURSOS",
        "ASPECTOS ECONÓMICOS REGIONALES": "ECONOMIA_Y_RECURSOS",
        "ECONOMÍA AGRÍCOLA": "ECONOMIA_Y_RECURSOS",
        "RECURSOS MINERALES": "ECONOMIA_Y_RECURSOS",
        "RECURSOS NATURALES": "ECONOMIA_Y_RECURSOS",
        "DEMOGRAFÍA RURAL": "POBLACION_Y_ETNIAS",
        "ETNIAS": "POBLACION_Y_ETNIAS",
        "GRUPOS ÉTNICOS": "POBLACION_Y_ETNIAS",
        "POBLACIÓN": "POBLACION_Y_ETNIAS",
        "POBLACIÓN ROM": "POBLACION_Y_ETNIAS",
    },
    "HISTORIA": {
        "": "GENERAL",
        "FIGURAS HISTÓRICAS": "GENERAL",
        "PERÍODO PRECOLOMBINO": "PRECOLOMBINO",
        "PRECOLOMBINO": "PRECOLOMBINO",
        "POBLAMIENTO": "PRECOLOMBINO",
        "CONQUISTA": "CONQUISTA",
        "LA CONQUISTA": "CONQUISTA",
        "COLONIA": "COLONIA",
        "LA COLONIA": "COLONIA",
        "REFORMAS BORBÓNICAS": "COLONIA",
        "ILUSTRACIÓN Y EXPEDICIÓN BOTÁNICA": "COLONIA",
        "INDEPENDENCIA": "INDEPENDENCIA",
        "LA INDEPENDENCIA": "INDEPENDENCIA",
        "GRAN COLOMBIA": "GRAN_COLOMBIA",
        "SIGLO XIX": "SIGLO_XIX",
        "SIGLO XIX - REPÚBLICA": "SIGLO_XIX",
        "OLIMPO RADICAL": "SIGLO_XIX",
        "ORGANIZACIÓN TERRITORIAL": "SIGLO_XIX",
        "PERÍODOS Y NOMBRES HISTÓRICOS": "SIGLO_XIX",
        "SIGLO XX": "SIGLO_XX",
        "EL BOGOTAZO": "SIGLO_XX",
        "HEGEMONÍA CONSERVADORA": "SIGLO_XX",
        "QUINQUENIO REYES": "SIGLO_XX",
        "REPÚBLICA LIBERAL": "SIGLO_XX",
        "VOTO FEMENINO": "SIGLO_XX",
        "MOVIMIENTOS POPULARES": "SIGLO_XX",
        "SALUD PÚBLICA": "SIGLO_XX",
        "ECONOMÍA": "SIGLO_XX",
        "CONSTITUCION DE 1991": "SIGLO_XX",
        "CONFLICTO ARMADO": "CONFLICTO_Y_PAZ",
        "GRUPOS ARMADOS": "CONFLICTO_Y_PAZ",
        "NARCOTRÁFICO": "CONFLICTO_Y_PAZ",
        "JUSTICIA TRANSICIONAL": "CONFLICTO_Y_PAZ",
        "POLÍTICA DE SEGURIDAD": "CONFLICTO_Y_PAZ",
    },
    "CONSTITUCION": {
        "": "GENERAL",
        "JERARQUÍA NORMATIVA": "PRINCIPIOS_FUNDAMENTALES",
        "PRINCIPIOS CONSTITUCIONALES": "PRINCIPIOS_FUNDAMENTALES",
        "PRINCIPIOS FUNDAMENTALES": "PRINCIPIOS_FUNDAMENTALES",
        "SOBERANÍA": "PRINCIPIOS_FUNDAMENTALES",
        "ASOCIACIÓN SINDICAL": "DERECHOS_Y_LIBERTADES",
        "DERECHO AL TRABAJO": "DERECHOS_Y_LIBERTADES",
        "DERECHOS FUNDAMENTALES": "DERECHOS_Y_LIBERTADES",
        "IGUALDAD": "DERECHOS_Y_LIBERTADES",
        "LIBERTAD DE CONCIENCIA": "DERECHOS_Y_LIBERTADES",
        "LIBERTAD DE LOCOMOCIÓN": "DERECHOS_Y_LIBERTADES",
        "LIBERTADES INDIVIDUALES": "DERECHOS_Y_LIBERTADES",
        "CONTROL CONSTITUCIONAL": "MECANISMOS_DE_PROTECCION",
        "MECANISMOS DE PROTECCIÓN": "MECANISMOS_DE_PROTECCION",
        "MECANISMOS DE PARTICIPACIÓN": "PARTICIPACION_Y_ELECCIONES",
        "ORGANIZACIÓN ELECTORAL": "PARTICIPACION_Y_ELECCIONES",
        "RAMA EJECUTIVA": "RAMAS_DEL_PODER_PUBLICO",
        "RAMA JUDICIAL": "RAMAS_DEL_PODER_PUBLICO",
        "RAMA LEGISLATIVA": "RAMAS_DEL_PODER_PUBLICO",
        "RAMAS DEL PODER PÚBLICO": "RAMAS_DEL_PODER_PUBLICO",
        "ORGANISMOS DE CONTROL": "ORGANISMOS_DE_CONTROL",
        "ÓRGANOS AUTÓNOMOS": "ORGANISMOS_DE_CONTROL",
        "DESCENTRALIZACIÓN": "ORGANIZACION_TERRITORIAL",
        "ENTIDADES TERRITORIALES": "ORGANIZACION_TERRITORIAL",
        "ESTADOS DE EXCEPCIÓN": "ESTADOS_DE_EXCEPCION",
        "REFORMAS CONSTITUCIONALES": "REFORMAS_CONSTITUCIONALES",
    },
}

KEY_PREFIX = {"CULTURA": "CUL", "GEOGRAFIA": "GEO", "HISTORIA": "HIS", "CONSTITUCION": "CON"}
FILE_NAME = {"CULTURA": "cultura", "GEOGRAFIA": "geografia", "HISTORIA": "historia", "CONSTITUCION": "constitucion"}


def normalize_text(text: str) -> str:
    """Normalization used only for duplicate detection."""
    text = unicodedata.normalize("NFC", text)
    return " ".join(text.split()).casefold()


def main() -> None:
    questions = json.loads(SOURCE.read_text(encoding="utf-8"))
    print(f"Loaded {len(questions)} questions from {SOURCE}")

    # Fail loudly on any subcategory the mapping does not cover
    unmapped = sorted({
        (q["category"], q["subcategory"])
        for q in questions
        if q["subcategory"] not in MAPPING.get(q["category"], {})
    })
    if unmapped:
        raise SystemExit(f"Unmapped (category, subcategory) pairs: {unmapped}")

    # Deduplicate by normalized text, preferring the copy with an explanation
    by_text: dict[str, dict] = {}
    order: list[str] = []
    for q in questions:
        norm = normalize_text(q["text"])
        if norm not in by_text:
            by_text[norm] = q
            order.append(norm)
        elif q.get("explanation") and not by_text[norm].get("explanation"):
            by_text[norm] = q
    deduped = [by_text[n] for n in order]
    print(f"Removed {len(questions) - len(deduped)} duplicate questions -> {len(deduped)} remain")

    # Group per category, remap subcategories and assign stable keys
    per_category: dict[str, list[dict]] = {code: [] for code in TAXONOMY}
    for q in deduped:
        per_category[q["category"]].append(q)

    QUESTIONS_DIR.mkdir(parents=True, exist_ok=True)
    used_subcats: dict[str, Counter] = {}
    for cat_code, cat_questions in per_category.items():
        out = []
        counts: Counter = Counter()
        for i, q in enumerate(cat_questions, start=1):
            sub_code = MAPPING[cat_code][q["subcategory"]]
            counts[sub_code] += 1
            # A few source questions have every choice at order 0; renumber them
            choices = q["choices"]
            orders = [c["order"] for c in choices]
            if len(set(orders)) != len(orders):
                choices = [{**c, "order": j} for j, c in enumerate(choices, start=1)]
                q = {**q, "choices": choices}
            out.append({
                "key": f"{KEY_PREFIX[cat_code]}-{i:04d}",
                "subcategory": sub_code,
                "text": q["text"].strip(),
                "difficulty": q["difficulty"],
                "points": q["points"],
                "hint": q.get("hint", ""),
                "explanation": q.get("explanation", ""),
                "choices": q["choices"],
            })
        used_subcats[cat_code] = counts
        path = QUESTIONS_DIR / f"{FILE_NAME[cat_code]}.json"
        path.write_text(json.dumps(out, ensure_ascii=False, indent=2) + "\n", encoding="utf-8")
        print(f"Wrote {len(out)} questions to {path} ({len(counts)} subcategories)")

    # Warn about taxonomy entries no question uses (kept anyway: they are valid targets)
    for cat_code, (_, subs) in TAXONOMY.items():
        unused = set(subs) - set(used_subcats[cat_code])
        if unused:
            print(f"Note: unused subcategories in {cat_code}: {sorted(unused)}")

    taxonomy_out = [
        {
            "code": cat_code,
            "name": cat_name,
            "subcategories": [{"code": c, "name": n} for c, n in subs.items()],
        }
        for cat_code, (cat_name, subs) in TAXONOMY.items()
    ]
    TAXONOMY_OUT.write_text(json.dumps(taxonomy_out, ensure_ascii=False, indent=2) + "\n", encoding="utf-8")
    print(f"Wrote taxonomy to {TAXONOMY_OUT}")


if __name__ == "__main__":
    main()

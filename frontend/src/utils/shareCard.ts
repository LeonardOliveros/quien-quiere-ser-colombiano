import type { GameResults } from '@/types'

// Paleta de la app (style.css) — el canvas no lee variables CSS
const COLORS = {
  bgTop: '#0B1B3A',
  bgBottom: '#12264F',
  card: 'rgba(255, 255, 255, 0.05)',
  cardBorder: 'rgba(255, 255, 255, 0.10)',
  yellow: '#FFCD00',
  blue: '#0033A0',
  red: '#CE1126',
  redLight: '#E8455B',
  emerald: '#10B981',
  text: '#FDF9EE',
  textDim: 'rgba(253, 249, 238, 0.6)'
}

const SIZE = 1080
const FONT = "'Segoe UI', 'Helvetica Neue', Arial, sans-serif"

function formatTime(seconds: number): string {
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  const secs = seconds % 60
  if (hours > 0) return `${hours}h ${minutes}m ${secs}s`
  if (minutes > 0) return `${minutes}m ${secs}s`
  return `${secs}s`
}

function roundRect(
  ctx: CanvasRenderingContext2D,
  x: number, y: number, w: number, h: number, r: number
) {
  ctx.beginPath()
  ctx.moveTo(x + r, y)
  ctx.arcTo(x + w, y, x + w, y + h, r)
  ctx.arcTo(x + w, y + h, x, y + h, r)
  ctx.arcTo(x, y + h, x, y, r)
  ctx.arcTo(x, y, x + w, y, r)
  ctx.closePath()
}

/** Dibuja la tarjeta de resultados (1080×1080) y la devuelve como PNG. */
export function renderShareCard(results: GameResults): Promise<Blob> {
  const canvas = document.createElement('canvas')
  canvas.width = SIZE
  canvas.height = SIZE
  const ctx = canvas.getContext('2d')!

  const pct = Math.round(results.percentage)
  const passed = results.percentage >= 70
  const accent = passed ? COLORS.emerald : COLORS.redLight
  const incorrect = results.total_questions - results.correct_answers

  // Fondo
  const bg = ctx.createLinearGradient(0, 0, SIZE, SIZE)
  bg.addColorStop(0, COLORS.bgTop)
  bg.addColorStop(1, COLORS.bgBottom)
  ctx.fillStyle = bg
  ctx.fillRect(0, 0, SIZE, SIZE)

  // Franja tricolor superior (proporciones de la bandera 2:1:1)
  ctx.fillStyle = COLORS.yellow
  ctx.fillRect(0, 0, SIZE, 24)
  ctx.fillStyle = COLORS.blue
  ctx.fillRect(0, 24, SIZE, 12)
  ctx.fillStyle = COLORS.red
  ctx.fillRect(0, 36, SIZE, 12)

  // Título
  ctx.textAlign = 'center'
  ctx.fillStyle = COLORS.yellow
  ctx.font = `800 52px ${FONT}`
  ctx.fillText('¿Quién Quiere Ser Colombiano?', SIZE / 2, 140)
  ctx.fillStyle = COLORS.textDim
  ctx.font = `400 30px ${FONT}`
  ctx.fillText('Resultados del examen', SIZE / 2, 190)

  // Anillo de puntaje
  const cx = SIZE / 2
  const cy = 430
  const radius = 165
  const ringWidth = 26

  ctx.lineWidth = ringWidth
  ctx.lineCap = 'round'
  ctx.strokeStyle = 'rgba(255, 255, 255, 0.12)'
  ctx.beginPath()
  ctx.arc(cx, cy, radius, 0, Math.PI * 2)
  ctx.stroke()

  if (pct > 0) {
    const start = -Math.PI / 2
    ctx.strokeStyle = accent
    ctx.beginPath()
    ctx.arc(cx, cy, radius, start, start + (Math.min(pct, 100) / 100) * Math.PI * 2)
    ctx.stroke()
  }

  ctx.fillStyle = COLORS.text
  ctx.font = `800 110px ${FONT}`
  ctx.fillText(`${pct}%`, cx, cy + 18)
  ctx.fillStyle = COLORS.textDim
  ctx.font = `400 32px ${FONT}`
  ctx.fillText('de acierto', cx, cy + 70)

  // Chip aprobado / no aprobado
  const chipText = passed ? '¡Aprobado!' : 'No aprobado'
  ctx.font = `700 40px ${FONT}`
  const chipW = ctx.measureText(chipText).width + 80
  const chipH = 72
  const chipY = 660
  roundRect(ctx, cx - chipW / 2, chipY, chipW, chipH, chipH / 2)
  ctx.fillStyle = passed ? 'rgba(16, 185, 129, 0.16)' : 'rgba(232, 69, 91, 0.16)'
  ctx.fill()
  ctx.lineWidth = 3
  ctx.strokeStyle = accent
  ctx.stroke()
  ctx.fillStyle = accent
  ctx.fillText(chipText, cx, chipY + 49)

  // Estadísticas
  const stats = [
    { value: String(results.correct_answers), label: 'Correctas', color: COLORS.emerald },
    { value: String(incorrect), label: 'Incorrectas', color: COLORS.redLight },
    { value: String(results.score), label: 'Puntos', color: COLORS.text },
    { value: formatTime(results.time_taken), label: 'Tiempo', color: COLORS.text }
  ]
  const tileW = 225
  const tileH = 130
  const gap = 20
  const totalW = stats.length * tileW + (stats.length - 1) * gap
  let x = (SIZE - totalW) / 2
  const tileY = 790

  for (const stat of stats) {
    roundRect(ctx, x, tileY, tileW, tileH, 18)
    ctx.fillStyle = COLORS.card
    ctx.fill()
    ctx.lineWidth = 2
    ctx.strokeStyle = COLORS.cardBorder
    ctx.stroke()

    ctx.fillStyle = stat.color
    ctx.font = `700 46px ${FONT}`
    ctx.fillText(stat.value, x + tileW / 2, tileY + 62)
    ctx.fillStyle = COLORS.textDim
    ctx.font = `400 26px ${FONT}`
    ctx.fillText(stat.label.toUpperCase(), x + tileW / 2, tileY + 102)
    x += tileW + gap
  }

  // Pie de página
  ctx.fillStyle = COLORS.textDim
  ctx.font = `400 28px ${FONT}`
  ctx.fillText(`Ponte a prueba en ${window.location.host}`, SIZE / 2, 1010)

  return new Promise((resolve, reject) => {
    canvas.toBlob(
      blob => (blob ? resolve(blob) : reject(new Error('No se pudo generar la imagen'))),
      'image/png'
    )
  })
}

/**
 * Comparte la tarjeta con la Web Share API (hoja nativa en móvil:
 * Instagram, LinkedIn, WhatsApp…). Si no está disponible, descarga el PNG.
 */
export async function shareResults(results: GameResults): Promise<'shared' | 'downloaded'> {
  const blob = await renderShareCard(results)
  const file = new File([blob], 'mis-resultados-colombiano.png', { type: 'image/png' })

  if (navigator.canShare?.({ files: [file] })) {
    try {
      await navigator.share({
        files: [file],
        title: '¿Quién Quiere Ser Colombiano?',
        text: `Obtuve ${Math.round(results.percentage)}% en el examen de ¿Quién Quiere Ser Colombiano? 🇨🇴`
      })
      return 'shared'
    } catch (err) {
      // Usuario canceló la hoja de compartir: no hacer nada más
      if ((err as DOMException).name === 'AbortError') return 'shared'
      // Otro fallo: caer a descarga
    }
  }

  const url = URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.href = url
  link.download = 'mis-resultados-colombiano.png'
  link.click()
  URL.revokeObjectURL(url)
  return 'downloaded'
}

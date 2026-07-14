package dynamodb

import (
	"fmt"
	"strconv"
)

// Key layout (single table, PK/SK strings, no GSIs):
//
//	Entity            PK                        SK
//	Counter           COUNTER#<entity>          META
//	Category          TAXONOMY                  CAT#<code>
//	SubCategory       TAXONOMY                  CAT#<cat>#SUB#<code>
//	Question          QBANK                     Q#<key>
//	User profile      USER#<id>                 PROFILE
//	Username unique   UNIQ#USERNAME#<username>  UNIQ (registered users only)
//	Auth token        TOKEN#<token>             META
//	Day rollup        DAY#<yyyy-mm-dd>          META
//	DAU dedup marker  DAY#<yyyy-mm-dd>          USER#<id>
//	GameSession       USER#<id>                 SESSION#<pad12(id)>
//	Session pointer   SESSION#<id>              META
//	GameAnswer        SESSION#<id>              ANSWER#<pad12(qid)>
//	QuestionHistory   SESSION#<id>              HISTORY#<pad12(qid)>
//	Recommendation    USER#<id>                 REC#<pad12(id)>
//
// Session/answer/recommendation IDs come from monotonically increasing
// counters, so the zero-padded SKs sort in creation order.
const (
	skMeta    = "META"
	skProfile = "PROFILE"
	skUniq    = "UNIQ"

	pkTaxonomy = "TAXONOMY"
	pkQBank    = "QBANK"

	prefixSession = "SESSION#"
	prefixAnswer  = "ANSWER#"
	prefixHistory = "HISTORY#"
	prefixRec     = "REC#"
	prefixCat     = "CAT#"
)

func pad12(id uint) string { return fmt.Sprintf("%012d", id) }

func pkCounter(entity string) string { return "COUNTER#" + entity }
func pkUser(id uint) string          { return "USER#" + strconv.FormatUint(uint64(id), 10) }
func pkUniqUsername(u string) string { return "UNIQ#USERNAME#" + u }
func pkToken(token string) string    { return "TOKEN#" + token }

// pkDay is the partition of one day's usage rollup (SK META) and its DAU
// dedup markers (SK USER#<id>). day is domain.MetricsDay formatted.
func pkDay(day string) string  { return "DAY#" + day }
func skDayUser(id uint) string { return "USER#" + strconv.FormatUint(uint64(id), 10) }

// pkSession is the partition holding the session pointer, answers and history.
func pkSession(id uint) string { return "SESSION#" + strconv.FormatUint(uint64(id), 10) }

// skSession is the session item's sort key under the user partition.
func skSession(id uint) string  { return prefixSession + pad12(id) }
func skAnswer(qid uint) string  { return prefixAnswer + pad12(qid) }
func skHistory(qid uint) string { return prefixHistory + pad12(qid) }
func skRec(id uint) string      { return prefixRec + pad12(id) }

func skCategory(code string) string { return prefixCat + code }
func skSubCategory(cat, sub string) string {
	return prefixCat + cat + "#SUB#" + sub
}
func skQuestion(key string) string { return "Q#" + key }

// Choice IDs are derived, not stored in a lookup table: questionID*100+order.
// This keeps them numeric (the frontend expects numbers), globally unique and
// stable across re-seeds. Orders are validated to stay below choiceIDFactor.
const choiceIDFactor = 100

func choiceID(questionID uint, order int) uint {
	return questionID*choiceIDFactor + uint(order)
}

func splitChoiceID(id uint) (questionID uint, order int) {
	return id / choiceIDFactor, int(id % choiceIDFactor)
}

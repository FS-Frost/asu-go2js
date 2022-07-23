package main

import (
	"github.com/FS-Frost/asu"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	globalAsu := map[string]interface{}{
		"newError":   asu.NewErrorJS,
		"newLine":    asu.NewLineJS,
		"newTime":    asu.NewTimeJS,
		"newTagB":    asu.NewTagBJS,
		"newTagPos":  asu.NewTagPosJS,
		"newTagMove": asu.NewTagMoveJS,
		"tagExists":  asu.TagExists,
		"replaceTag": asu.ReplaceTag,
	}

	js.Global.Set("asu", globalAsu)
}

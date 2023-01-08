package ediff_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cqroot/ediff"
)

func TestDiff(t *testing.T) {
	ed := ediff.New("vim")
	ed.SetEditorArgs([]string{
		"-N", "-n",
		"-c", "%s/item/obj/",
		"-c", "wq",
	})
	ed.SetIgnoreEditorError(true)

	ed.AppendItems([]string{
		"item 1",
		"item 2",
		"item 3",
	})

	pairs, err := ed.Diff()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []ediff.DiffPair{
		{Prev: "item 1", Curr: "obj 1"},
		{Prev: "item 2", Curr: "obj 2"},
		{Prev: "item 3", Curr: "obj 3"},
	}, pairs)
}

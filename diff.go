package ediff

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

func (e Ediff) runEditor(file string) error {
	args := append(e.editorArgs, file)

	edcmd := exec.Command(e.editor, args...)
	edcmd.Stdin = os.Stdin
	edcmd.Stdout = os.Stdout
	err := edcmd.Run()

	if e.ignoreEditorError {
		return nil
	} else {
		return err
	}
}

func (e Ediff) Diff() ([]DiffPair, error) {
	// Create temp file
	tmp, err := os.CreateTemp("", "ediff-")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmp.Name())

	// Write items to temp file
	if _, err = tmp.WriteString(strings.Join(e.items, "\n")); err != nil {
		return nil, err
	}
	if err := tmp.Sync(); err != nil {
		return nil, err
	}

	// Run editor
	if err := e.runEditor(tmp.Name()); err != nil {
		return nil, err
	}

	// Read new items from temp file
	if _, err := tmp.Seek(0, 0); err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(tmp)
	scanner.Split(bufio.ScanLines)

	idx := 0
	pairs := make([]DiffPair, 0)
	for scanner.Scan() {
		newItem := scanner.Text()
		if newItem == e.items[idx] {
			idx += 1
			continue
		}

		pairs = append(pairs, DiffPair{
			Prev: e.items[idx],
			Curr: newItem,
		})
		idx += 1
	}

	return pairs, nil
}

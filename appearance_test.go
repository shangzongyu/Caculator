package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/egoist/mygo"
	"github.com/shangzongyu/Caculator/locales"
)

func TestRememberAppearance(t *testing.T) {
	const dataDirEnv = "CACULATOR_TEST_APPEARANCE_DIR"
	const phaseEnv = "CACULATOR_TEST_APPEARANCE_PHASE"
	if dir := os.Getenv(dataDirEnv); dir != "" {
		mygo.App.SetPath(mygo.PathUserData, dir)
		app := calculatorView{appearance: savedAppearance()}
		if os.Getenv(phaseEnv) == "save" {
			if got := appearances[app.appearance].id; got != "warm-sand" {
				t.Fatalf("without preferences, appearance = %q, want warm-sand", got)
			}
			if err := app.selectAppearance(1); err != nil {
				t.Fatal(err)
			}
			return
		}
		menu := app.appearanceMenu(locales.In("en"), func(int) {})
		if got := appearances[app.appearance].id; got != "graphite" || !menu.Submenu[app.appearance].IsChecked() {
			t.Fatalf("after restart, appearance = %q, want graphite selected", got)
		}
		return
	}

	// Each process opens the same test directory with a fresh preference cache.
	dir := t.TempDir()
	binary, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	for _, phase := range []string{"save", "restore"} {
		cmd := exec.CommandContext(t.Context(), binary, "-test.run=^TestRememberAppearance$")
		cmd.Env = append(os.Environ(), dataDirEnv+"="+dir, phaseEnv+"="+phase)
		if output, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("%s preferences: %v\n%s", phase, err, output)
		}
	}
}

package iflag_test

import (
	"os"
	"testing"
	"time"

	"github.com/itozll/iflag"
)

func TestBoolFlag(t *testing.T) {
	os.Setenv("BF", "true")
	bf := iflag.NewBoolFlag(&iflag.Bool{
		Name:         "test",
		Aliases:      []string{"t"},
		DefaultValue: false,
		EnvVars:      []string{"bf"},
	})

	bf.PreExec()
	if !*bf.Destination {
		t.Fatal("need true")
	}

	bf.SetWithString("false", iflag.ChangeLevelMin)
	if !*bf.Destination {
		t.Fatal("need true")
	}

	bf.SetWithString("false", iflag.ChangeLevelEnv)
	if *bf.Destination {
		t.Fatal("need false")
	}

	bf.SetWithString("true", iflag.ChangeLevelConfig)
	if !*bf.Destination {
		t.Fatal("need true")
	}

	bf.SetWithString("false", iflag.ChangeLevelEnv)
	if !*bf.Destination {
		t.Fatal("need true")
	}
}

func TestStringSliceFlag(t *testing.T) {
	v := iflag.NewStringSliceFlag(&iflag.StringSlice{
		Name:         "foo",
		DefaultValue: []string{"hello", "world"},
	})

	v.PreExec()

	if len(v.Value()) != 2 {
		t.Fatal("len must be equal 2")
	}

	v.Append(",")
	t.Log(v.Value())
}

func TestIntFlag(t *testing.T) {
	v := iflag.NewIntFlag(&iflag.Int{
		Name:         "foo",
		DefaultValue: 10,
	})

	v.PreExec()

}

func TestDurationFlag(t *testing.T) {
	v := iflag.NewDurationFlag(&iflag.Duration{
		Name:         "foo",
		DefaultValue: 1 * time.Minute,
	})

	v.PreExec()

	v.SetWithString("1h20m", iflag.ChangeLevelConfig)
	t.Log(v.Value())
	v.Append("2h10s")
	t.Log(v.Value())
}

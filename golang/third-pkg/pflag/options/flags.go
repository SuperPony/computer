package options

import (
	"flag"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/spf13/pflag"
)

type NamedFlagSets struct {
	Order    []string
	FlagSets map[string]*pflag.FlagSet
}

var once sync.Once

var FlagGroup *NamedFlagSets

func Init() {
	once.Do(func() {
		FlagGroup = &NamedFlagSets{
			Order:    make([]string, 0),
			FlagSets: make(map[string]*pflag.FlagSet),
		}
		NamedOptions = NewOptions()
		NamedOptions.AddFlags()
	})
}

func (nfs *NamedFlagSets) FlagSet(name string) *pflag.FlagSet {
	if _, isSet := nfs.FlagSets[name]; !isSet {
		nfs.FlagSets[name] = pflag.NewFlagSet(name, pflag.ExitOnError)
		nfs.Order = append(nfs.Order, name)
	}

	return nfs.FlagSets[name]
}

// WordSepNormalizeFunc 用于将非标准的标志位转换为标准格式
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

func InitFlags(flags *pflag.FlagSet) {
	flags.SetNormalizeFunc(WordSepNormalizeFunc)
	flags.AddGoFlagSet(flag.CommandLine)
}

func Parse() {
	if len(os.Args) <= 1 {
		return
	}

	var name string
	for _, v := range FlagGroup.Order {
		if v == os.Args[1] {
			name = v
			break
		}
	}

	if name == "" {
		panic("flag is not found")
	}

	InitFlags(FlagGroup.FlagSets[name])

	if err := FlagGroup.FlagSets[name].Parse(os.Args[1:]); err != nil {
		log.Fatalln(err.Error())
	}
}
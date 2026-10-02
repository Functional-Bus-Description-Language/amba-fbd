package csync

import (
	_ "embed"
	"log"
	"os"
	"path"
	"sync"
	"text/template"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/pkg"

	"github.com/Functional-Bus-Description-Language/amba-fbd/internal/args"
	"github.com/Functional-Bus-Description-Language/amba-fbd/internal/c"
	"github.com/Functional-Bus-Description-Language/amba-fbd/internal/utils"
)

var busWidth int64

var readType c.Type

//go:embed templates/amba_fbd.h
var headerTmplStr string
var headerTmpl = template.Must(template.New("C-Sync amba_fbd.h").Parse(headerTmplStr))

type headerFormatters struct {
	BusWidth int64
}

func Generate(bus *fn.Block, pkgsConsts map[string]*pkg.Package) {
	busWidth = bus.Width

	err := os.MkdirAll(args.CSync.Path, os.FileMode(int(0775)))
	if err != nil {
		log.Fatalf("generate C-Sync: %v", err)
	}

	hFile, err := os.Create(path.Join(args.CSync.Path, "amba_fbd.h"))
	if err != nil {
		log.Fatalf("generate C-Sync: %v", err)
	}

	readType = c.WidthToReadType(bus.Width)

	hFmts := headerFormatters{
		BusWidth: bus.Width,
	}

	err = headerTmpl.Execute(hFile, hFmts)
	if err != nil {
		log.Fatalf("generate C-Sync: %v", err)
	}

	blocks := utils.CollectBlocks(bus, nil, []string{})
	utils.ResolveBlockNameConflicts(blocks)

	var wg sync.WaitGroup
	defer wg.Wait()

	for _, b := range blocks {
		wg.Add(1)
		go genBlock(b, &wg)
	}

	err = hFile.Close()
	if err != nil {
		log.Fatalf("generate C-Sync: %v", err)
	}

	if args.CSync.LinuxMmapIface {
		GenLinuxMmapIface(bus)
	}
}

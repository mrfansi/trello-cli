package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7D56F4"))
	cellStyle   = lipgloss.NewStyle().Padding(0, 1)
)

type Renderer struct {
	JSON bool
	W    io.Writer
}

func New(jsonMode bool) *Renderer {
	return &Renderer{JSON: jsonMode, W: os.Stdout}
}

func (r *Renderer) Raw(v any) error {
	enc := json.NewEncoder(r.W)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// Table prints headers + rows. JSON mode falls back to objects keyed by header.
func (r *Renderer) Table(headers []string, rows [][]string) error {
	if r.JSON {
		out := make([]map[string]string, 0, len(rows))
		for _, row := range rows {
			obj := map[string]string{}
			for i, h := range headers {
				if i < len(row) {
					obj[h] = row[i]
				}
			}
			out = append(out, obj)
		}
		return r.Raw(out)
	}
	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("240"))).
		Headers(headers...).
		Rows(rows...).
		StyleFunc(func(row, _ int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}
			return cellStyle
		})
	fmt.Fprintln(r.W, t)
	return nil
}

func (r *Renderer) KeyValue(pairs [][2]string) error {
	if r.JSON {
		obj := map[string]string{}
		for _, p := range pairs {
			obj[p[0]] = p[1]
		}
		return r.Raw(obj)
	}
	maxKey := 0
	for _, p := range pairs {
		if l := len(p[0]); l > maxKey {
			maxKey = l
		}
	}
	for _, p := range pairs {
		fmt.Fprintf(r.W, "%-*s  %s\n", maxKey, p[0], p[1])
	}
	return nil
}

func Truncate(s string, n int) string {
	s = strings.ReplaceAll(s, "\n", " ")
	if len(s) <= n {
		return s
	}
	if n <= 1 {
		return s[:n]
	}
	return s[:n-1] + "…"
}

func Deref[T any](p *T) T {
	var zero T
	if p == nil {
		return zero
	}
	return *p
}

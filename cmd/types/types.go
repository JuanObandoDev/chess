package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Southclaws/supervillain"
	"github.com/sanpezlo/chess/internal/api/auth/discord"
	"github.com/sanpezlo/chess/internal/api/auth/github"
	"github.com/sanpezlo/chess/internal/resources/user"
	"github.com/sanpezlo/chess/internal/web"
)

func generate(name, prefix string, objs ...interface{}) {
	output := strings.Builder{}
	output.WriteString(`import * as z from "zod"

`)

	for _, v := range objs {
		output.WriteString(supervillain.StructToZodSchemaWithPrefix(prefix, v))
	}

	err := os.WriteFile(
		fmt.Sprintf("web/types/%s.ts", name),
		[]byte(output.String()),
		os.ModePerm,
	)
	if err != nil {
		panic(err)
	}
}

func main() {
	generate("Error", "API", web.Error{})
	generate("User", "", user.User{})

	generate("GitHub", "", github.Link{}, github.Callback{})
	generate("Discord", "", discord.Link{}, discord.Callback{})
}

#
# Makefile for running pandoc on all Markdown docs ending in .md
#
PROJECT = fdx

MD_PAGES = $(shell ls -1 *.md)

HTML_PAGES = $(shell ls -1 *.md | sed -E 's/.md/.html/g')

MAN_PAGES = fdx2txt txt2fdx

MAN_HTML = fdx2txt.html txt2fdx.html

build: $(HTML_PAGES) $(MD_PAGES) $(MAN_HTML) LICENSE.html

$(HTML_PAGES): $(MD_PAGES) .FORCE
	pandoc --metadata title=$(basename $@) -s --to html5 $(basename $@).md -o $(basename $@).html \
	    --template=page.tmpl
	@if [ $@ = "README.html" ]; then mv README.html index.html; fi

$(MAN_HTML): .FORCE
	pandoc docs/$(basename $@).md -s --to html5 -o $@ \
		--template=page.tmpl

LICENSE.html: LICENSE
	pandoc --metadata title="$(PROJECT) License" -s --from Markdown --to html5 LICENSE -o license.html \
	    --template=page.tmpl

clean:
	@if [ -f index.html ]; then rm *.html; fi
	#@if [ -f docs/index.html ]; then rm docs/*.html; fi

.FORCE:

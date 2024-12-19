# make new FOLDER=graphs FILE=1591-strange-printer-ii

new:
	cp template/template.py "${FOLDER}/${FILE}.py"
	cp template/template.go "${FOLDER}/${FILE}.go"
	cp template/template_test.go "${FOLDER}/${FILE}_test.go"

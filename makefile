toolsGo := tools.go
toolsDir := bin/tools
toolPkgs := $(shell go list -f '{{join .Imports " "}}' ${toolsGo})
toolCmds := $(foreach tool,$(notdir ${toolPkgs}),${toolsDir}/${tool})
$(foreach cmd,${toolCmds},$(eval $(notdir ${cmd})Cmd := ${cmd}))

go.mod: ${toolsGo}
	go mod tidy
	touch go.mod

${toolCmds}: go.mod
	go build -o $@ $(filter %/$(@F),${toolPkgs})


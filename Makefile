.PHONY: init
init:
ifeq ($(shell uname -s),Darwin)
	@grep -r -l slack-reminder * .goreleaser.yml | xargs sed -i "" "s/go-cli-template/$$(basename `git rev-parse --show-toplevel`)/"
else
	@grep -r -l slack-reminder * .goreleaser.yml | xargs sed -i "s/go-cli-template/$$(basename `git rev-parse --show-toplevel`)/"
endif

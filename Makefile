# git remote add github git@github.com:wit-go/log.git

redomod:
	rm -f go.*
	unset GO111MODULES && go mod init
	unset GO111MODULES && go mod tidy

github:
	git push origin master
	git push origin devel
	git push origin --tags
	git push github master
	git push github devel
	git push github --tags
	@echo
	@echo check https://github.com/wit-go/log
	@echo

init-github:
	git push -u github master
	git push -u github devel
	git push github --tags

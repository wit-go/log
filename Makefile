# git remote add github git@github.com:wit-go/log.git

github:
	git push origin master
	git push origin devel
	git push origin --tags
	git push github master
	git push github devel
	git push github --tags
	@echo
	@echo check https://github.com/wit-go/gui
	@echo

init-github:
	git push -u github master
	git push -u github devel
	git push github --tags

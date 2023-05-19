#!/bin/bash
#

WORKING_BRANCH=$(git branch | grep -E "\* " | cut -d \  -f 2)
if [ "$WORKING_BRANCH" = "gh-pages" ]; then
	git commit -am "publishing to gh-pages branch"
	git push origin gh-pages
else
	echo "You're in $WORKING_BRANCH branch"
	echo "You need to pull in changes to the gh-pages branch to publish"
	read -p "process Y/n " YES_NO
	if [ "$YES_NO" = "Y" ] || [ "$YES_NO" = "y" ]; then
		echo "Committing and pushing to $WORKING_BRANCH"
		git commit -am "commiting to $WORKING_BRANCH"
		git push origin "$WORKING_BRANCH"
		echo "Changing branchs from $WORKING_BRANCH to gh-pages"
		git checkout gh-pages
		echo "Merging changes from origin main"
		git pull origin "$WORKING_BRANCH"
		git commit -am "Merge from $WORKING_BRANCH to gh-pages"
		echo "Building website in gh-pages branch"
		make installer.sh man
		make -f website.mak
		echo "Commiting changes after website/installer.sh build"
		git commit -am "Merging site updates in gh-pages"
		echo "Pushing changes up to gh-pages"
		git push origin gh-pages
		echo "Changing back to your working branch $WORKING_BRANCH"
		git checkout "$WORKING_BRANCH"
	fi
fi

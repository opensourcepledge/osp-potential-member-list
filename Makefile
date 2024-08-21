.PHONY: publish

publish:
	rsync -Pvrthl --delete --exclude .git --info=progress2 ./ yavin:/srv/www/osp-potential-member-list

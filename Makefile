# © 2024 Vlad-Stefan Harbuz <vlad@vlad.website>
#
# SPDX-License-Identifier: Apache-2.0

.PHONY: publish

bin/osp-potential-member-list:
	go build -o bin/ .

publish: bin/osp-potential-member-list
	rsync -Pvrthl --delete --exclude .git --info=progress2 ./ yavin:/srv/www/osp-potential-member-list

#!/bin/bash
TS=`date +"%Y-%m-%dT%H:%M:%S"`

curl 'https://www.cdc.gov.tw/RSS/RssXml/Hh094B49-DRwe2RR4eFfrQ?type=1' > downloads/cdc_rss.xml
go run tw_cdc_rss_to_news.go > ../data/news.json


if [[ `git status --porcelain ../data/news.json` ]]; then
    # Changes
    echo 'RSS Updated.'
    git add '../data/news.json';
    git commit -S -m "Update nesw.json in ${TS}";
    # git push;

else
	echo "There are no news. ${TS}"
fi





#!/bin/bash
TS=`date +"%Y-%m-%dT%H:%M:%S"`

curl "https://sheets.googleapis.com/v4/spreadsheets/1qh20J-5rGVIEjLcGKJnfj7huAp-nCxsd-fJdmh3yZKY/values/%E6%AA%A2%E9%A9%97%E4%BA%BA%E6%95%B8"'!'"A1:L?majorDimension=ROWS&key=${GOOGLE_API_KEY}" > downloads/inspections_summary_raw.json
#curl 'https://raw.githubusercontent.com/kelvin2go/covid19-tw/master/latest.json' > downloads/patients_raw.json
go run tw_leaflu_google_sheet_json_to_inspections_summary_data.go > ../data/inspections_summary.json


DIFF_RESULT=`git diff --numstat ../data/inspections_summary.json | awk '{print $1}'`
if [[ $DIFF_RESULT -gt 2 ]]; then
    # Changes
    echo 'LeafLu inspections summary data has been updated.'
    git add '../data/inspections_summary.json';
    git commit -S -m "Update inspections_summary.json in ${TS}";
    # git push;

else
	echo "There are no new inspections summary data. ${TS}"
fi





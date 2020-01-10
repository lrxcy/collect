# Authority Jim.Weng
# Test Env MBP

#!/bin/bash

# PATH=/Users/mac/go/src/github.com/jimweng/ATourOfGO/basic
cd /Users/mac/go/src/github.com/jimweng/ATourOfGO/basic
# save the git log as temp variable
lastCommit=`git log|head -5 | tail -1`
echo $lastCommit

# back to your folder where you put your modified needed file
cd -


DATE=`date "+%Y-%m-%d"`
echo $DATE

# | status | used day | task description |
# | --- | --- |
# | done,defer,ongoing,abort |  | (Phabricater Task number if available) descript of what you have done |
# | statusjim | 1 day | task_to_descrip |
# everything woudl be add at after 26 line
# E.g.(add in line 8):
# sed -i ".bak" '8i\
# 8 This is Line 8' jimweng_2017Q4.md

sed -i ".bak" '26i\
| done | 1 day | '  $lastcommit  '|' jimweng_2017Q4.md



# sed -i ".bak" '27i\
# ' jimweng_2017Q4.md

# if need to change file content
# sed -i ".test" "s@statusjim@done@g" jimweng_2017Q4.md
# sed -i ".test" "s@task_to_descrip@Nodescription@g" jimweng_2017Q4.md

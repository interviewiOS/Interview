#!/bin/sh

file_path="${1}"
webhook_key="${2}"


# 上传一个文件到企业微信
# path - 上传的文件路径
# bwGurlToken - 调用接口凭证, 企业微信机器人webhookurl中的key参数
# return - 文件的media_id
function uploadFileBusinessWeixin()
{
    local path="${1}"
    local bwGurlToken="${2}"

    local sendResult=$(curl -k -fsSL -H "Content-Type: multipart/form-data" -X POST "https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?key=${bwGurlToken}&type=file" --form "upload=@${path}")
    grep -oE "\"media_id\":[^\"]*\"[^\"]+" <<<"${sendResult}" | awk -F '"' '{print $4}'
}

# 发送文件到企业微信群
# 依赖 - getAtUserList
# path - 发送的文件路径
# bwGurl - 企业微信机器人webhookurl
# atUser - 需要被@的用户ID
function sendFileToBusinessWeixinGroup()
{
    local path="${1}"
    local bwGurl="${2}"
    local atUser="${3}"

    local bwGurlToken="$(echo "${bwGurl}" | grep -oE "key=[a-z0-9-]*" | awk -F '=' '{print $2}')"
    local mediaId="$(uploadFileBusinessWeixin "${path}" "${bwGurlToken}")"
    local sendResult=$(curl -k -fsSL -H "Content-Type:application/json" -X POST "${bwGurl}" -d "{\"msgtype\":\"file\",\"file\":{\"media_id\":\"${mediaId}\",\"mentioned_list\":[\"${atUser}\"]}}")
    echoImpl "将\"${path}\"上传，得到media_id为\"${mediaId}\"的文件，发送到机器人\"${bwGurlToken}\"所在的企业微信：${sendResult}"
}

function test() {
    local path="testUpload.csv"
    local bwGurlToken="c4a79450-c10c-4549-907a-d746c220e286"
    local bwGurl="https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=c4a79450-c10c-4549-907a-d746c220e286"
    local mediaId="$(uploadFileBusinessWeixin "${path}" "${bwGurlToken}")"
    local sendResult=$(curl -k -fsSL -H "Content-Type:application/json" -X POST "${bwGurl}" -d "{\"msgtype\":\"file\",\"file\":{\"media_id\":\"${mediaId}\",\"mentioned_list\":[\"${atUser}\"]}}")
    echo "将\"${path}\"上传，得到media_id为\"${mediaId}\"的文件，发送到机器人\"${bwGurlToken}\"所在的企业微信：${sendResult}"
}

test
# ==================================================================
# google-chrome
# ------------------------------------------------------------------ 
MYENVPATH=/root/conf
apt-get update -y && apt-get install curl -y

apt -y update \
    && apt install -y gpg-agent \
    && curl -LO https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
    && (dpkg -i ./google-chrome-stable_current_amd64.deb || apt-get install -fy) \
    && curl -sSL https://dl.google.com/linux/linux_signing_key.pub | apt-key add \
    && rm google-chrome-stable_current_amd64.deb \
    && rm -rf /var/lib/apt/lists/* 


sed -e '/chrome/ s/^#*/#/' -i /opt/google/chrome/google-chrome
echo 'exec -a "$0" "$HERE/chrome" "$@" --user-data-dir="$HOME/.config/chrome" --no-sandbox --disable-dev-shm-usage' >> /opt/google/chrome/google-chrome

cp ${MYENVPATH}/google-chrome.desktop /usr/share/applications/

# ==================================================================
# microsoft-edge
# ------------------------------------------------------------------ 
#apt -y update 
#apt install software-properties-common apt-transport-https wget ca-certificates gnupg2 ubuntu-keyring -y
#wget -O- https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor | tee /usr/share/keyrings/microsoft-edge.gpg
#echo 'deb [signed-by=/usr/share/keyrings/microsoft-edge.gpg] https://packages.microsoft.com/repos/edge stable main' | tee /etc/apt/sources.list.d/microsoft-edge.list
#apt update
#apt install microsoft-edge-stable -y
##cp  /root/conf/microsoft-edge.desktop /usr/share/applications/

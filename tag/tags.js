tags = new Array();
document.querySelectorAll('#all-topic-tags>a').forEach(function (e) {
    /tag\/(\S+?)\//.test(e.href);
    tags.push({
        "Name": e.title,
        "Slug": RegExp.$1,
        "TranslatedName": e.innerText.replace(/\s/g, '')
    });
});
newWindow = window.open('', 'frame_name');
newWindow.document.write(JSON.stringify(tags));
newWindow.document.close();

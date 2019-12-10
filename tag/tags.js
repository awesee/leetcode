tags = [];
document.querySelectorAll('#all-topic-tags>a').forEach(function (e) {
    tags.push({
        "Name": e.title || e.firstElementChild.innerText.trim(),
        "Slug": /tag\/(\S+?)\//.exec(e.href)[1],
        "TranslatedName": e.innerText.trim().replace(/\n/g, ''),
    });
});
newWindow = window.open('', 'frame_name');
newWindow.document.write(JSON.stringify(tags));
newWindow.document.close();

tags = new Array();
$("#all-topic-tags>a").each(function () {
    /tag\/(\S+)\//.test(this.href);
    tags.push({"Slug": RegExp.$1});
});
document.write(JSON.stringify(tags));

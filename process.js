const fs = require('fs');
const args = require('yargs').argv;

let rawdata = fs.readFileSync(args.source);
let posts = JSON.parse(rawdata);
console.log(posts.length);

const outputDir = `data-processed`
if (!fs.existsSync(outputDir)) {
    fs.mkdirSync(outputDir);
}

// source: https://tecadmin.net/how-to-parse-command-line-arguments-in-nodejs/
if (args.cmd === 'comment_likes_sum') {
    posts.sort(compareByCommentLikes)
    savePosts(`cls`, args.source, posts)
} else if (args.cmd === 'likes_total') {
    posts.sort(compareByLikesTotal)
    savePosts(`lt`, args.source, posts)
} else if (args.cmd === 'lopt2c') { // likes_op_plus_top_2_comments
    // TODO: implement
}

function compareByCommentLikes(postA, postB) {
    let postACommentLikes = commentLikes(postA);
    let postBCommentLikes = commentLikes(postB);
    return postBCommentLikes - postACommentLikes;
}

function commentLikes(post) {
    let likeSum = 0;
    for (let i = 0; i < post.comments.length; i++) {
        likeSum += post.comments[i].likes;
    }
    return likeSum;
}

function compareByLikesTotal(postA, postB) {
    let postALikes = likesTotal(postA);
    let postBLikes = likesTotal(postB);
    return postBLikes - postALikes;
}

function likesTotal(post) {
    return post.likes + commentLikes(post);
}

function savePosts(prefix, source, posts) {
    const outputDir = `data-processed/${prefix}`
    if (!fs.existsSync(outputDir)) {
        fs.mkdirSync(outputDir);
    }
    const sourceFilename = source.split('/')[1];
    const outputFilename = `${outputDir}/${sourceFilename}`;

    fs.writeFile(outputFilename, JSON.stringify(posts, null, 4), function(err) {
        if(err) {
            console.log(err);
        } else {
            console.log("JSON saved to " + outputFilename);
        }
    });
}
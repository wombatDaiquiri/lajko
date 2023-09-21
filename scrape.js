const axios = require('axios');
const cheerio = require('cheerio');
const fs = require("fs");

let posts = [];

const sleepTime = 10;

if (!fs.existsSync('data-snapshots')) {
    fs.mkdirSync('data-snapshots');
}

async function main() {
    for (let i= 0; i < 1; i++) {
        const pageNumber = i+1;
        console.log('scraping page ' + pageNumber);

        let resp;
        try {
            // this will later be replaced by our own database so treat it as debugging data
            resp = await axios.get(`https://www.hejto.pl/gorace/okres/6h/strona/${pageNumber}`);
            // resp = await axios.get(`https://www.hejto.pl/najnowsze/strona/${pageNumber}`);
            await sleepBetweenRequests();
        } catch (err) {
            console.log(err);
            continue;
        }

        const postURLs = extractPostURLs(resp);

        for (let j= 0; j < postURLs.length; j++) {
            console.log(`scraping ${postURLs[j]} START`);
            await scrapeMicroblogPost(postURLs[j]);
            await sleepBetweenRequests();
            console.log(`scraping ${postURLs[j]} END`);
        }

        console.log('page ' + pageNumber + ' scraped')
    }
    console.log('all pages scraped')
    savePosts();
}

main().
    then(() => console.log('posts length in main promise:' + posts.length)).
    catch(err => console.log(err));

// ------------ only function declarations from this point on -----------

// -------------- this is our end goal - to populate posts array with posts to be saved as JSON ----------------
function savePosts() {
    // source: https://stackoverflow.com/questions/2573521/how-do-i-output-an-iso-8601-formatted-string-in-javascript
    let date = new Date();
    const outputFilename = `data-snapshots/${date.toISOString()}.json`;

    fs.writeFile(outputFilename, JSON.stringify(posts, null, 4), function(err) {
        if(err) {
            console.log(err);
        } else {
            console.log("JSON saved to " + outputFilename);
        }
    });
}

// -------------- this is our "business logic" for scraping posts (without HTTP requests) ----------------

function extractPostURLs(pageResponse) {
    const $ = cheerio.load(pageResponse.data);
    const articles = $('article');
    const postURLs = [];
    articles.each((i, el) => {
        const postPath = $(el).find('div.items-start a').last().attr('href');
        console.log(postPath)

        if (postPath) {
            // microblog
            postURLs.push(`https://www.hejto.pl${postPath}`);
        } else {
            // TODO: article
        }
    })
    return postURLs;
}

function scrapeMicroblogPost(postURL) {
    return axios.get(postURL).then(function (response) {
        const $ = cheerio.load(response.data);
        const article = $('article');

        const post = {images: []};

        const originalPost = article.find('div').first();
        originalPost.find('div').first().children().each((i, el) => {
            if (i === 0) {
                // header
                post.authorURL = $(el).find('a').first().attr('href');
                post.author = post.authorURL.split('/')[2];
                post.avatar = $(el).find('img').first().attr('src');
                post.likes = $(el).find('button').last().text();
                post.url = postURL
            } else
            if (i === 1) {
                // content
                // post.content = $(el).text()
                post.content = $(el).find('.parsed.text-sm').html()
            } else {
                // whatever, but we'll just add all images to post.images
                $(el).find('img').each((i, el) => {
                    post.images.push($(el).attr("src"))
                })
            }
        })

        post.comments = [];

        article.children().each((i, el) => {
            if (i===0) {
                // original post
            }
            if (i===1) {
                // comments
                $(el).find('div.relative.flex.flex-col.gap-2.pt-px').children().each((i, el) => {
                    if ($(el).hasClass('gap-2')) {
                        // list of subcomments

                        $(el).children().each((i, el) => {
                            // subcomment
                            // append to subcomments array of last comment
                            const subcomment = {};

                            subcomment.authorURL = $(el).find('a').first().attr('href');
                            subcomment.author = subcomment.authorURL.split('/')[2];
                            subcomment.avatar = $(el).find('img').first().attr('src');
                            subcomment.likes = $(el).find('button').last().text();

                            // subcomment.content = $(el).find('div.parsed.text-sm').text();
                            subcomment.content = $(el).find('div.parsed.text-sm').html();
                            subcomment.images = [];

                            $(el).find('div.w-full img').each((i, el) => {
                                subcomment.images.push($(el).attr("src"))
                            })

                            post.comments[post.comments.length - 1].subcomments.push(subcomment);
                        })
                    } else {
                        // top-level comment
                        console.log(i)
                        console.log($(el).text())
                        console.log($(el).find('div.parsed.text-sm').text())

                        const comment = {subcomments: []};

                        comment.authorURL = $(el).find('a').first().attr('href');
                        comment.author = comment.authorURL.split('/')[2];
                        comment.avatar = $(el).find('img').first().attr('src');
                        comment.likes = $(el).find('button').last().text();

                        // comment.content = $(el).find('div.parsed.text-sm').text();
                        comment.content = $(el).find('div.parsed.text-sm').html();
                        comment.images = [];

                        $(el).find('div.w-full img').each((i, el) => {
                            comment.images.push($(el).attr("src"))
                        })

                        post.comments.push(comment);
                    }
                })
            }
        })

        console.log(post);
        posts.push(post);

        //savePosts();
    }).catch(err => console.log(err));
}

// -------------- utils ----------------

// sleepBetweenRequests is used to rate-limit requests
//
// at this point in time I'm not sure if this is correct or not
async function sleepBetweenRequests()  { await delay(sleepTime) }

// source: https://masteringjs.io/tutorials/node/sleep
function delay(time) {
    return new Promise(resolve => setTimeout(resolve, time));
}

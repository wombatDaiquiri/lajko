const axios = require('axios');
const cheerio = require('cheerio');
const fs = require("fs");

let posts = [];

const sleepTime = 10;

if (!fs.existsSync('data-snapshots')) {
    fs.mkdirSync('data-snapshots');
}

async function main() {
    for (let i= 0; i < 20; i++) {
        const pageNumber = i+1;
        console.log('scraping page ' + pageNumber);

        let resp;
        try {
            resp = await axios.get(`https://www.hejto.pl/najnowsze/strona/${pageNumber}`);
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

        const post = {}

        const originalPost = article.find('div').first();
        originalPost.find('div').first().children().each((i, el) => {
            if (i === 0) {
                // header
                post.authorURL = $(el).find('a').first().attr('href');
                post.author = post.authorURL.split('/')[2];
                post.avatar = $(el).find('img').first().attr('src');
                post.likes = $(el).find('button').last().text();
            }
            if (i === 1) {
                // content
                post.content = $(el).text()
            }
        })

        post.comments = [];

        article.children().each((i, el) => {
            if (i===0) {
                // original post
            }
            if (i===1) {
                // comments
                $(el).children().first().find('div.bg-grey-250').each((i, el) => {
                    console.log(i)
                    console.log($(el).text())
                    console.log($(el).find('div.parsed.text-sm').text())

                    const comment = {};

                    comment.authorURL = $(el).find('a').first().attr('href');
                    comment.author = comment.authorURL.split('/')[2];
                    comment.avatar = $(el).find('img').first().attr('src');
                    comment.likes = $(el).find('button').last().text();

                    comment.content = $(el).find('div.parsed.text-sm').text();
                    comment.images = [];

                    $(el).find('div.w-full img').each((i, el) => {
                        comment.images.push($(el).attr("src"))
                    })

                    post.comments.push(comment);
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

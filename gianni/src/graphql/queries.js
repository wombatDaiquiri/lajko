import { gql } from 'graphql-request';

export const GET_POSTS = gql`
    query GetPosts {
        posts {
            author {
                username
            }
            title
            content
        }    
    }
`;

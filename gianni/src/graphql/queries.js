import { gql } from 'graphql-request';

export const GET_POSTS_BY_CATEGORY = gql`
    query GetPostsByCategory($category: String!) {
        posts(query: {orderBy: $category, limit: 5}) {
            author {
                username
            }
            title
            content
        }
    }
`;
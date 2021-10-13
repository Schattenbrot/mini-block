import PostListItem from './PostListItem';

const PostList = (props) => {
  const { posts } = props;

  return (
    <div>
      <h1>All Posts</h1>
      <ul>
        {posts.map((post) => (
          <PostListItem key={post._id} post={post} />
        ))}
      </ul>
    </div>
  );
};

export default PostList;

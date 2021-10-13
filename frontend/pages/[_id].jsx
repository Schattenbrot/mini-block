import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { deletePost, getOnePost } from '../utils/db-utils';

const PostPage = () => {
  const [post, setPost] = useState({});
  const router = useRouter();
  const { _id } = router.query;

  const backHandler = () => {
    router.push('http://localhost:3000/');
  };

  const deleteHandler = () => {
    deletePost(_id);
    router.replace('http://localhost:3000/');
  };

  useEffect(() => {
    if (_id) {
      getOnePost(_id).then((post) => setPost(post));
    }
  }, [_id]);

  if (!post) {
    return <p>Loading...</p>;
  }

  return (
    <div>
      <h1>{post.title}</h1>
      <p>{post.text}</p>
      <p>{post.updated_at}</p>
      <p>{post.created_at}</p>

      <div>
        <button onClick={backHandler}>View All Posts</button>
        <button onClick={deleteHandler}>Delete</button>
      </div>
    </div>
  );
};

export default PostPage;

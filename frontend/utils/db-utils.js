export const getAllPosts = async () => {
  const promise = await fetch('http://localhost:4000/v1/posts');
  const json = await promise.json();
  const posts = json.posts;

  return posts;
};

export const getOnePost = async (_id) => {
  const promise = await fetch(`http://localhost:4000/v1/posts/${_id}`);
  const json = await promise.json();
  const post = json.post;

  return post;
};

export const deletePost = async (_id) => {
  fetch(`http://localhost:4000/v1/posts/${_id}`, {
    method: 'DELETE',
  });
};

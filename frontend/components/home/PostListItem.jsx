const PostListItem = (props) => {
  const { post } = props;
  const { _id, title, text, updated_at, created_at } = post;

  const deleteHandler = () => {
    fetch(`http://localhost:4000/v1/posts/${_id}`, {
      method: 'DELETE',
    });
  };

  return (
    <li onClick={deleteHandler}>
      <h2>{title}</h2>
      <p>{text}</p>
      <div>
        <p>{created_at}</p>
        <p>{updated_at}</p>
      </div>
    </li>
  );
};

export default PostListItem;

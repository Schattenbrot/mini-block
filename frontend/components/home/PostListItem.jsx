import { useRouter } from 'next/router';

const PostListItem = (props) => {
  const { post } = props;
  const { _id, title, text, created_at } = post;
  const router = useRouter();

  const showPostHandler = () => {
    router.push(`http://localhost:3000/${_id}`);
  };

  return (
    <li onClick={showPostHandler}>
      <h2>{title}</h2>
      <p>{text}</p>
      <p>{created_at}</p>
    </li>
  );
};

export default PostListItem;

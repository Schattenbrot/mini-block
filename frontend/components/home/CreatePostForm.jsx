import { useState } from 'react';

const CreatePostForm = () => {
  const [isExpanded, setIsExpanded] = useState(false);
  const [title, setTitle] = useState('');
  const [text, setText] = useState('');

  const titleInputHandler = (event) => {
    setTitle(event.target.value);
  };

  const textInputHandler = (event) => {
    setText(event.target.value);
  };

  const expandHandler = () => {
    setIsExpanded((prevState) => !prevState);
  };

  const submitHandler = (event) => {
    event.preventDefault();

    fetch('http://localhost:4000/v1/posts', {
      method: 'POST',
      header: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        title,
        text,
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.response.ok) {
          setTitle('');
          setText('');

          alert('success');
        } else {
          alert('something failed');
        }
      });
  };

  if (!isExpanded) {
    return (
      <div>
        <button type='button' onClick={expandHandler}>
          Add Post
        </button>
      </div>
    );
  }

  return (
    <form onSubmit={submitHandler}>
      <div>
        <label htmlFor='title'>Title:</label>
        <input
          type='text'
          name='title'
          id='title'
          value={title}
          onInput={titleInputHandler}
        />
      </div>
      <div>
        <label htmlFor='Text'>Text:</label>
        <input
          type='text'
          name='Text'
          id='Text'
          value={text}
          onInput={textInputHandler}
        />
      </div>
      <div>
        <button type='button' onClick={expandHandler}>
          Close
        </button>
        <button type='submit'>Add Post</button>
      </div>
    </form>
  );
};

export default CreatePostForm;

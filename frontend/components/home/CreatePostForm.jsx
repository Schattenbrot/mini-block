import { useState } from 'react';
import useInput from '../../hooks/use-input';

const CreatePostForm = () => {
  const [isExpanded, setIsExpanded] = useState(false);
  const {
    value: title,
    hasError: titleHasError,
    isValid: titleIsValid,
    changeHandler: titleChangeHandler,
    resetHandler: titleResetHandler,
    blurHandler: titleBlurHandler,
  } = useInput((value) => value.trim() !== '');
  const {
    value: text,
    hasError: textHasError,
    isValid: textIsValid,
    changeHandler: textChangeHandler,
    resetHandler: textResetHandler,
    blurHandler: textBlurHandler,
  } = useInput((value) => value.trim() !== '');

  const formIsValid = titleIsValid && textIsValid;

  const expandHandler = () => {
    setIsExpanded((prevState) => !prevState);
  };

  const submitHandler = (event) => {
    event.preventDefault();

    if (!formIsValid) {
      return;
    }

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
          titleResetHandler();
          textResetHandler();

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
          onInput={titleChangeHandler}
          onBlur={titleBlurHandler}
        />
      </div>
      <div>
        <label htmlFor='Text'>Text:</label>
        <input
          type='text'
          name='Text'
          id='Text'
          value={text}
          onInput={textChangeHandler}
          onBlur={textBlurHandler}
        />
      </div>
      <div>
        <button type='button' onClick={expandHandler}>
          Close
        </button>
        <button type='submit' disabled={!formIsValid}>
          Add Post
        </button>
      </div>
    </form>
  );
};

export default CreatePostForm;

import { useState } from 'react';

const useInput = (validateValue) => {
  const [enteredValue, setEnteredValue] = useState('');
  const [isTouched, setIsTouched] = useState(false);

  const valueIsValid = validateValue(enteredValue);
  const hasError = !valueIsValid && isTouched;

  const changeHandler = (event) => {
    setEnteredValue(event.target.value);
  };

  const blurHandler = (event) => {
    setIsTouched(true);
  };

  const resetHandler = () => {
    setEnteredValue('');
  };

  return {
    value: enteredValue,
    hasError,
    isValid: valueIsValid,
    changeHandler,
    blurHandler,
    resetHandler,
  };
};

export default useInput;

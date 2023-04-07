import React, { ChangeEvent, useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch } from '../../store';
import { selectURL } from '../../features/NewBookmark/selects';
import { updateURL } from '../../features/NewBookmark/slice';
import createBookmark from '../../features/NewBookmark/createBookmark';

const NewBookmark = () => {
  const dispatch = useDispatch<AppDispatch>();

  const urlValue = useSelector(selectURL);
  const urlChanged = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      dispatch(updateURL(evt.target.value));
    },
    [dispatch],
  );

  const formSubmitted = useCallback(() => {
    dispatch(createBookmark(urlValue));
  }, [dispatch, urlValue]);

  return (
    <div className="uk-container">
      <div className="uk-margin-top">
        <input
          className="uk-input uk-form-large uk-border-rounded"
          type="text"
          value={urlValue}
          onChange={urlChanged}
          placeholder="Link"
          // eslint-disable-next-line jsx-a11y/no-autofocus
          autoFocus
        />
        <div className="uk-flex uk-flex-row-reverse uk-margin-small-top">
          <button
            type="button"
            onClick={formSubmitted}
            className="uk-button uk-button-primary uk-border-rounded"
          >
            Add
          </button>
        </div>
      </div>
    </div>
  );
};

export default NewBookmark;

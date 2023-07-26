import React, { ChangeEvent, useCallback } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Button, Container, Flex, TextInput } from '@mantine/core';
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
    <Container>
      <Flex my="md">
        <TextInput
          sx={{ flex: 1 }}
          value={urlValue}
          onChange={urlChanged}
          placeholder="Link"
          // eslint-disable-next-line jsx-a11y/no-autofocus
          autoFocus
        />
        <Button variant="outline" onClick={formSubmitted} ml="sm">
          Add
        </Button>
      </Flex>
    </Container>
  );
};

export default NewBookmark;

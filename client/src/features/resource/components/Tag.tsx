import React, { useEffect, useRef, useState } from 'react';
import { PlusOutlined } from '@ant-design/icons';
import type { InputProps } from 'antd';
import { Flex, Input, Tag, theme, Tooltip, Form, InputRef } from 'antd';

const InputTag: React.FC<InputProps> = ({ id, value, ...props }) => {
  const { token } = theme.useToken();
  const form = Form.useFormInstance();
  const [tags, setTags] = useState<string[]>(value as string[]);
  const [inputVisible, setInputVisible] = useState(false);
  const [inputValue, setInputValue] = useState('');
  const [editInputIndex, setEditInputIndex] = useState(-1);
  const [editInputValue, setEditInputValue] = useState('');

  const inputRef = useRef<InputRef>(null);
  const editInputRef = useRef<InputRef>(null);

  useEffect(() => {
    if (inputVisible) {
      inputRef.current?.focus();
    }
  }, [inputVisible]);

  useEffect(() => {
    editInputRef.current?.focus();
  }, [editInputValue]);

  const handleClose = (removedTag: string) => {
    const newTags = tags.filter((tag) => tag !== removedTag);

    if (inputValue.split(':')[1] != undefined) {
      setTags(newTags);
      if (id) {
        form.setFieldsValue({ [id]: newTags });
      }
    }
  };

  const showInput = () => {
    setInputVisible(true);
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  const handleInputConfirm = () => {
    if (inputValue.split(':')[1] != undefined) {
      if (inputValue && !tags.includes(inputValue)) {
        if (inputValue) setTags([...tags, inputValue]);

        if (id) {
          form.setFieldsValue({ [id]: [...tags, inputValue] });
        }
      }
    } else {
      if (id) {
        form.setFields([
          {
            name: [id],
            errors: ['ohh fuck'],
          },
        ]);
      }
    }

    setInputVisible(false);
    setInputValue('');
  };

  const handleEditInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEditInputValue(e.target.value);
  };

  const handleEditInputConfirm = () => {
    const newTags = [...tags];
    newTags[editInputIndex] = editInputValue;

    if (inputValue.split(':')[1] != undefined) {
      setTags(newTags);
      if (id) {
        form.setFieldsValue({ [id]: newTags });
      }
    }
    setEditInputIndex(-1);
    setEditInputValue('');
  };

  const tagPlusStyle: React.CSSProperties = {
    height: 22,
    background: token.colorBgContainer,
    borderStyle: 'dashed',
  };

  const tagInputStyle: React.CSSProperties = {
    width: 64,
    height: 22,
    marginInlineEnd: 8,
    verticalAlign: 'middle',
  };

  return (
    <Flex gap="4px 0" wrap>
      {tags.map<React.ReactNode>((tag, index) => {
        if (editInputIndex === index) {
          return (
            <Input
              ref={editInputRef}
              key={tag}
              size="small"
              style={tagInputStyle}
              value={editInputValue}
              onChange={handleEditInputChange}
              onBlur={handleEditInputConfirm}
              onPressEnter={handleEditInputConfirm}
            />
          );
        }

        const isLongTag = tag.length > 15;

        const tagElem = (
          <Tag key={tag} closable={true} style={{ userSelect: 'none' }} onClose={() => handleClose(tag)}>
            <span
              onDoubleClick={(e) => {
                if (index !== 0) {
                  setEditInputIndex(index);
                  setEditInputValue(tag);
                  e.preventDefault();
                }
              }}
            >
              {isLongTag ? `${tag.slice(0, 15)}...` : tag}
            </span>
          </Tag>
        );
        return isLongTag ? (
          <Tooltip title={tag} key={tag}>
            {tagElem}
          </Tooltip>
        ) : (
          tagElem
        );
      })}

      {inputVisible ? (
        <Input
          ref={inputRef}
          type="text"
          size="small"
          style={tagInputStyle}
          value={inputValue}
          onChange={handleInputChange}
          onBlur={handleInputConfirm}
          onPressEnter={handleInputConfirm}
        />
      ) : (
        <Tag style={tagPlusStyle} icon={<PlusOutlined />} onClick={showInput}>
          New Tag
        </Tag>
      )}
    </Flex>
  );
};

export default InputTag;

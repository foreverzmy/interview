import React, { FC, useRef, memo, useLayoutEffect, useState } from 'react';
import * as CodeMirror from 'codemirror';

import './style.scss';
import 'codemirror/lib/codemirror.css'; // 核心样式文件
import 'codemirror/theme/dracula.css'; // 主题文件
import 'codemirror/addon/display/fullscreen.css'; // 全屏样式文件

interface ICodeMirrorProps {
  value: string;
  onChange: (text: string) => void;
}

const Coder: FC<ICodeMirrorProps> = ({ value, onChange }) => {
  const ref = useRef<HTMLTextAreaElement>(null);
  const editor = useRef<CodeMirror.EditorFromTextArea>(null)
  const [isSetValue, setIsSetValue] = useState(false);

  useLayoutEffect(() => {
    if (editor.current === null) {
      editor.current = CodeMirror.fromTextArea(ref.current as HTMLTextAreaElement, {
        mode: 'Markdown', // HMTL混合模式
        indentUnit: 4,  // 缩进单位为4
        lineNumbers: false,	//显示行号
        theme: "dracula",	//设置主题
        lineWrapping: true,	//代码折叠
        styleActiveLine: true, // 当前行背景高亮
        foldGutter: true,
        autofocus: true,
        dragDrop: false,
        gutters: ["CodeMirror-linenumbers", "CodeMirror-foldgutter"],
      })

      editor.current.on('change', () => {
        onChange(editor.current.getValue())
      })
    }

    if (value !== '' && !isSetValue) {
      editor.current.setValue(value);
      setIsSetValue(true);
    }

  }, [value])

  return <textarea ref={ref} ></textarea >
}

export default memo(Coder);

import React, { FC, useMemo, useCallback } from 'react';
import './style.scss';

export interface IPaginationProps {
  current: number;
  size: number;
  total: number;
  onChange: (current: number, size: number) => void;
}

const Pagination: FC<IPaginationProps> = ({ current, size, total, onChange }) => {
  const totalPage = useMemo(() =>
    total / size === 0 ? total / size : Math.ceil(total / size
    ), [total, size])

  // 总页数
  const pages = useMemo(() => {
    if (totalPage < 5) {
      return [...Array(totalPage).keys()].map(i => i + 1);
    }

    if (current < 5) {
      return [...Array(5).keys()].map(i => i + 1);
    }

    if (current > totalPage - 3) {
      return [...Array(5).keys()].map(i => i + totalPage - 4);;
    }

    return [...Array(5).keys()].map(i => i + current - 2);;
  }, [size, total]);

  const handleChange = useCallback((next: number) => {
    onChange(next, size);
  }, []);

  return (
    <section className="pagination">
      每页 {size} 条，共 {total} 条 &nbsp;
      <button disabled={current < 3} onClick={() => handleChange(current - 1)}>&lt;</button>
      {pages.length == 5 && pages[0] > 1 && (
        <>
          <button>1</button>
          <span className="pagination__ellipsis">•••</span>
        </>
      )}
      {
        pages.map(page => {
          return (<button key={page} onClick={() => handleChange(page)}>{page}</button>)
        })
      }
      {pages.length == 5 && pages[4] < totalPage && (
        <>
          <span className="pagination__ellipsis">•••</span>
          <button >{totalPage}</button>
        </>
      )}
      <button disabled={current + 2 > totalPage} onClick={() => handleChange(current + 1)}>&gt;</button>
    </section>
  )
}

export default Pagination;

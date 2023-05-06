import React from 'react';

const FilterBar = () => (
  <div>
    <input type="text" className="uk-input" />
    <select>
      <option>All</option>
    </select>
    <button type="button">Filter</button>
  </div>
);

export default FilterBar;

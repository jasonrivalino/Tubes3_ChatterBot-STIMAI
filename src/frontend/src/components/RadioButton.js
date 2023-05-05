import React, { useState } from 'react';
import '../styles/RadioButton.css';

function RadioButton(props) {
  const [selectedValue, setSelectedValue] = useState(props.defaultValue);

  function handleChange(event) {
    setSelectedValue(event.target.value);
    props.onChange(event.target.value); // Memanggil fungsi onChange pada props
  }

  return (
    <div className='radioButton'>
      <label>
        <input
          type="radio"
          name={props.name}
          value={props.value1}
          checked={selectedValue === props.value1}
          onChange={handleChange}
        />
        {props.label1}
      </label>

      <br/><br/>
      
      <label>
        <input
          type="radio"
          name={props.name}
          value={props.value2}
          checked={selectedValue === props.value2}
          onChange={handleChange}
        />
        {props.label2}
      </label>
    </div>
  );
}

export default RadioButton;
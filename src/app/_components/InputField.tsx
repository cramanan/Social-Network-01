import React from "react";

interface InputFieldProps {
  label: string;
  type: string;
  name: string;
  id: string;
}

const InputField: React.FC<InputFieldProps> = ({ label, type, name }) => {
  return (
    <div>
      <input
        type={type}
        name={name}
        id={name}
        className="w-[350px] px-4 py-3.5 rounded-xl border border-white bg-transparent text-white text-xl justify-start items-center gap-2.5 inline-flex mb-4 placeholder-white"
        placeholder={label}
      />
    </div>
  );
};

export default InputField;

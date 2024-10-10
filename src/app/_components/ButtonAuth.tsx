import React from "react";

interface ButtonProps {
  label: string;
  onClick?: () => void;
}

const ButtonAuth: React.FC<ButtonProps> = ({ label, onClick }) => {
  return (
    <button
      className="w-full md:w-[200px] h-[46px] hover:bg-gradient-to-tr from-[#c5b8fa] via-[#9f79e9] to-[#e0d3ea] border border-bg-gradient-to-tr from-[#4821f9] via-[#6f46c0] to-[#e0d3ea] rounded-xl flex justify-center items-center"
      onClick={onClick}
    >
      <div className="text-white text-xl font-semibold font-['Noto Sans']">
        {label}
      </div>
    </button>
  );
};

export default ButtonAuth;

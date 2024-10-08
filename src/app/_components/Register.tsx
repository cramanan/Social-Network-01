import { Smartphone } from "./icons/smartphone";

export const Register = () => {
  return (
    <div className="w-screen h-screen flex justify-center items-center">
      <div className="w-[1450px] h-[2100px] bg-gradient-to-bl from-[#1667e0] to-[#e492e5] rounded-tr-[20px] rounded-br-[20px] shadow-2xl border backdrop-blur-[53px]"></div>
      <div className="w-[1550px] h-[2100px] bg-white flex justify-center items-center rounded-tl-[20px] rounded-bl-[20px] shadow-2xl borderbackdrop-blur-[53px] ">
        <Smartphone />
      </div>
    </div>
  );
};

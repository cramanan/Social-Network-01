import { Smartphone } from "./icons/smartphone";
import InputField from "./InputField";
import ButtonAuth from "./ButtonAuth";

export const Register = () => {
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="flex flex-col md:flex-row w-full max-w-6xl h-full md:h-[748px] ">
        <div className="w-full md:w-1/2 h-1/2 md:h-full relative bg-gradient-to-bl from-[#1667e0] to-[#e492e5] rounded-tl-[20px] rounded-bl-[20px] shadow border border-white backdrop-blur-[53px] flex flex-col justify-start items-center p-4 md:p-8">
          <div className="flex flex-col md:gap-10 w-full">
            <div className="text-white  text-4xl font-semibold font-['Noto Sans']">
              Register
            </div>
            <div className="flex flex-col justify-center items-center">
              <InputField label="Email" type="text" name="email" id="email" />
              <InputField
                label="Password"
                type="password"
                name="password "
                id="password"
              />
              <InputField
                label="First Name"
                type="text"
                name="firstname"
                id="firstname"
              />
              <InputField
                label="Last Name"
                type="text"
                name="lastname"
                id="lastname"
              />
              <InputField
                label="Nickname"
                type="text"
                name="nickname"
                id="nickname"
              />
              <InputField
                label="Date of Birth"
                type="date"
                name="dob"
                id="dob"
              />

              <div className="w-full flex flex-col items-center gap-2">
                <ButtonAuth label="Sign up" />
                <div className="flex items-center  w-full">
                  <div className="flex-grow border-t border-[#4c4c4c]"></div>
                  <div className="text-[#4c4c4c] text-base font-medium font-['Noto Sans']">
                    Or
                  </div>
                  <div className="flex-grow border-t border-[#4c4c4c]"></div>
                </div>
                <div className="text-center text-white text-base font-medium font-['Noto Sans']">
                  You have an account?
                </div>
                <ButtonAuth label="Sign in" />
              </div>
            </div>
          </div>
        </div>
        <div className="w-full md:w-1/2 h-1/2 md:h-full relative bg-white/90 rounded-tr-[20px] rounded-br-[20px] shadow border border-white backdrop-blur-[53px] flex items-center justify-center">
          <div className="w-full h-full rounded-r-2xl flex items-center justify-center">
            <Smartphone />
          </div>
        </div>
      </div>
    </div>
  );
};

import React from "react";

const Section = ({ children, className = "" }) => (

  <div className={`flex-col justify-start items-start ${className} flex`}>
    {children}
  </div>
);

const InputField = ({ label }) => (

  <div className="w-[400px] px-4 py-3.5 rounded-xl border border-white justify-start items-center gap-2.5 inline-flex">
    <div className="text-white text-xl font-normal font-['Noto Sans']">
      {label}
    </div>
  </div>
);

const Button = ({ label, className = "" }) => (

  <div
    className={`w-[400px] px-2.5 py-3.5 bg-gradient-to-tr from-[#4821f9] via-[#6f46c0] to-[#e0d3ea] rounded-xl justify-center items-center gap-2.5 inline-flex ${className}`}
  >
    <div className="text-white text-xl font-semibold font-['Noto Sans']">
      {label}
    </div>
  </div>
);

export const Login = () => {
return (

<div className="justify-items-center bg-ffffff">
<div className="w-[460px] h-[748px] px-[30px] bg-gradient-to-bl from-[#1667e0] to-[#e492e5] rounded-tr-[20px] rounded-br-[20px] shadow border border-white backdrop-blur-[53px] flex-col justify-start items-start inline-flex">
<div className="self-stretch h-[30px]" />
<div className="h-[675px] pt-[19.50px] pb-[31.50px] flex-col justify-start items-start gap-[46px] flex">
<Section className="self-stretch gap-[29px] inline-flex">
<Section className="gap-3.5">
<Section>
<div className="text-white text-4xl font-semibold font-['Noto Sans']">
Login
</div>
</Section>
<Section className="gap-[25px]">
<InputField label="Username" />
<Section className="gap-3">
<InputField label="Password" />
<div className="justify-start items-center gap-1 inline-flex">
<div className="w-[18px] h-[18px] relative" />
<div className="text-white text-base font-medium font-['Noto Sans']">
Remember me
</div>
</div>
</Section>
<Section className="justify-center items-center gap-3">
<div className="text-white text-base font-medium font-['Noto Sans']">
Forgot password ?
</div>
</Section>
</Section>
</Section>
<Button label="login" />
<div className="justify-start items-center gap-5 inline-flex">
<div className="w-[170px] h-[0px] border-2 border-[#4c4c4c]" />
<div className="text-[#4c4c4c] text-base font-medium font-['Noto Sans']">
Or
</div>
<div className="w-[170px] h-[0px] border-2 border-[#4c4c4c]" />
</div>
<div className="w-[239px] text-center text-white text-base font-medium font-['Noto Sans']">
Don’t have an account ? Signup
</div>
<Button label="Signup" className="h-[54px]" />
</Section>
<div className="self-stretch px-1.5 py-1 bg-gradient-to-b from-[#616161] to-[#616161] rounded-md justify-between items-center inline-flex">
<div className="justify-start items-start gap-2.5 flex">
<div className="text-white text-base font-normal font-['Noto Sans']">
Terms & Conditions
</div>
</div>
<div className="justify-start items-start gap-2.5 flex">
<div className="text-white text-base font-normal font-['Noto Sans']">
Support
</div>
</div>
<div className="justify-start items-start gap-2.5 flex">
<div className="text-white text-base font-normal font-['Noto Sans']">
Customer Care
</div>
</div>
</div>
</div>
</div>
</div>
);
};

"use client";

import React, { useState } from "react";
import { ProfileCircle } from "./ProfileCircle";

export const NewPost = () => {
  const [postText, setPostText] = useState("");
  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleTextChange = (e: React.ChangeEvent<HTMLTextAreaElement>) =>
    setPostText(e.target.value);

  const toggleModal = () => setIsModalOpen(!isModalOpen);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // Logique de soumission du post
    toggleModal();
  };

  return (
    <div className={isModalOpen ? "modal-open" : ""}>
      <div>
        <button
          className="px-4 py-2 relative bg-slate-50  hover:bg-slate-300 p-4 rounded-lg shadow-lg flex flex-row items-center gap-4"
          onClick={toggleModal}
        >
          <ProfileCircle />
          <div className="text-slate-500  font-extralight font-['Inter']">
            Create your Post
          </div>
        </button>
      </div>
      {isModalOpen && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-70 backdrop-blur-sm h-full  z-50">
          <div className=" relative border border-white bg-gradient-to-tr from-[#9ac0fa] to-[#efc0f0d7] p-6 rounded-lg shadow-lg w-1/2">
            <ProfileCircle className="absolute  top-4 left-9 " />
            <h2 className="text-l text-white font-semibold mb-4 flex justify-center ">
              Create your Post
            </h2>

            <textarea
              className="w-full  rounded-xl border border-gray-300  text-white text-xl justify-start items-center gap-2.5 inline-flex mb-4 placeholder-white"
              value={postText}
              onChange={handleTextChange}
              placeholder="What's on your mind?"
            />
            <div className="flex justify-end">
              <button
                className="px-4 py-2 bg-gray-500 text-white rounded mr-2"
                onClick={toggleModal}
              >
                Cancel
              </button>
              <button
                className="px-4 py-2 bg-blue-500 text-white rounded"
                type="submit"
                onClick={handleSubmit}
              >
                Submit
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

import { type FC, type FormEvent, useRef } from "react";
import type Todo from "~/types/Todo";

type Props = {
  onSubmit: (todo: Todo) => void;
};

const CreateTodoForm: FC<Props> = ({ onSubmit }) => {
  const titleRef = useRef<HTMLInputElement | null>(null);
  const descriptionRef = useRef<HTMLTextAreaElement | null>(null);

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!titleRef.current || !descriptionRef.current) return;

    onSubmit({
      title: titleRef.current.value,
      description: descriptionRef.current.value,
      status: false,
    });

    titleRef.current.value = "";
    descriptionRef.current.value = "";
  };

  return (
    <form
      className="relative w-64 rounded-md border border-gray-700 bg-gray-800 p-6 shadow hover:bg-gray-700"
      onSubmit={handleSubmit}
    >
      <div className="group relative z-0 mb-5 w-full">
        <input
          type="title"
          name="title"
          id="title"
          className="peer block w-full appearance-none border-0 border-b-2 border-gray-600 bg-transparent px-0 py-2.5  text-sm text-white focus:border-blue-500 focus:outline-none focus:ring-0"
          placeholder=""
          required
          ref={titleRef}
        />
        <label
          htmlFor="title"
          className="absolute top-3 -z-10 origin-[0] -translate-y-6 scale-75 transform text-sm text-gray-400 duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:start-0 peer-focus:-translate-y-6 peer-focus:scale-75 peer-focus:font-medium peer-focus:text-blue-500 rtl:peer-focus:left-auto rtl:peer-focus:translate-x-1/4"
        >
          Title
        </label>
      </div>
      <div className="mb-5">
        <label
          htmlFor="message"
          className="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
        >
          Description
        </label>
        <textarea
          id="description"
          className="block w-full rounded-lg border border-gray-600 bg-gray-700 p-2.5 text-sm text-white placeholder-gray-400 focus:border-blue-500 focus:ring-blue-500"
          required
          ref={descriptionRef}
        ></textarea>
      </div>
      <button
        type="submit"
        className="rounded-lg bg-blue-600 px-5 py-2.5 text-center text-sm font-medium  text-white hover:bg-blue-700 focus:outline-none focus:ring-4 focus:ring-blue-800"
      >
        Create
      </button>
    </form>
  );
};

export default CreateTodoForm;

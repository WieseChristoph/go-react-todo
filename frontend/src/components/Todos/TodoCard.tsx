import { type FC } from "react";
import type Todo from "~/types/Todo";

type Props = {
  todo: Todo;
  onClick: () => void;
  onDelete: () => void;
};

const TodoCard: FC<Props> = ({ todo, onClick, onDelete }) => {
  const handleDelete = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    // prevent onClick from firing
    e.stopPropagation();
    onDelete();
  };

  return (
    <div
      className="relative w-64 rounded-md border border-gray-700 bg-gray-800 p-6 shadow hover:bg-gray-700"
      onClick={onClick}
    >
      <div className="absolute left-0 top-0 p-2">
        <div
          className={`h-[10px] w-[10px] rounded-full ${
            todo.status ? "bg-green-600" : "bg-red-600"
          }`}
        ></div>
      </div>
      <div className="absolute right-0 top-0 p-2">
        <button
          className="m-0 p-0 text-gray-400 hover:text-gray-500"
          onClick={handleDelete}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            height={20}
            width={20}
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          >
            <path d="M18 6 6 18" />
            <path d="m6 6 12 12" />
          </svg>
        </button>
      </div>
      <span className="text-xs text-gray-500">
        {todo.created_at ? new Date(todo.created_at).toDateString() : "-"}
      </span>
      <div className="break-words text-2xl font-bold text-white">
        {todo.title}
      </div>
      <p className="break-words text-gray-400">{todo.description}</p>
    </div>
  );
};

export default TodoCard;

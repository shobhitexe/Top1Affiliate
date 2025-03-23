"use client";

import {
  ColumnDef,
  flexRender,
  getCoreRowModel,
  getPaginationRowModel,
  useReactTable,
} from "@tanstack/react-table";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Button } from "./button";

interface DataTableProps<TData, TValue> {
  columns: ColumnDef<TData, TValue>[];
  data: TData[];
}

export function DataTable<TData, TValue>({
  columns,
  data,
}: DataTableProps<TData, TValue>) {
  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
  });

  return (
    <div className="rounded-md overflow-auto">
      <Table className="min-w-full w-full relative">
        <TableHeader>
          {table.getHeaderGroups().map((headerGroup) => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map((header) => {
                return (
                  <TableHead key={header.id}>
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                          header.column.columnDef.header,
                          header.getContext()
                        )}
                  </TableHead>
                );
              })}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows?.length ? (
            table.getRowModel().rows.map((row) => (
              <TableRow
                key={row.id}
                data-state={row.getIsSelected() && "selected"}
              >
                {row.getVisibleCells().map((cell) => (
                  <TableCell key={cell.id}>
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="h-24 text-center">
                No results.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>

      <div className="flex items-center justify-between py-5">
        <div className="relative flex items-center justify-between">
          <select
            name="pagination"
            id="pagination"
            size={1}
            className="space-x-2 cursor-pointer p-2 rounded-lg bg-white border border-gray-300 appearance-none pr-8"
            onChange={(e) => table.setPageSize(Number(e.target.value))}
          >
            <option value="10" className="cursor-pointer">
              10
            </option>
            <option value="20" className="cursor-pointer">
              20
            </option>
            <option value="40" className="cursor-pointer">
              40
            </option>
            <option value="50" className="cursor-pointer">
              50
            </option>
          </select>
          <div className="absolute right-4 top-1/2 transform -translate-y-1/2 pointer-events-none">
            ▼
          </div>
        </div>

        <div className="flex items-center justify-end space-x-2 gap-2 self-end">
          <Button
            onClick={() => table.previousPage()}
            disabled={!table.getCanPreviousPage()}
            size={"icon"}
            className="cursor-pointer"
          >
            {"<"}
          </Button>
          <div>
            {table.getState().pagination.pageIndex + 1} / {table.getPageCount()}
          </div>
          <Button
            onClick={() => table.nextPage()}
            disabled={!table.getCanNextPage()}
            size={"icon"}
            className="cursor-pointer"
          >
            {">"}
          </Button>
        </div>
      </div>
    </div>
  );
}

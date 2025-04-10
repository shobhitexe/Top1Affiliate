"use client";

import { useState } from "react";
import { ChevronDown, ChevronRight, User } from "lucide-react";
import {
  Card,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import DetailsDropdown from "./DetailsDropdown";

interface Affiliate {
  id: string;
  affiliateId: string;
  name: string;
  commission: number;
  country: string;
  earnings: number;
  recruits: number;
  children?: Affiliate[];
}

interface TreeNodeProps {
  affiliate: Affiliate;
  level: number;
}

const TreeNode = ({ affiliate, level }: TreeNodeProps) => {
  const [expanded, setExpanded] = useState(level < 1);
  const hasChildren = affiliate.children && affiliate.children.length > 0;

  return (
    <div className="mb-2">
      <div className="flex items-start">
        {hasChildren && (
          <Button
            variant="ghost"
            size="sm"
            className="h-7 w-7 p-0 mr-1"
            onClick={() => setExpanded(!expanded)}
          >
            {expanded ? <ChevronDown size={14} /> : <ChevronRight size={14} />}
          </Button>
        )}
        {!hasChildren && <div className="w-7 mr-1" />}

        <Card className="w-full max-w-xs border border-black/10">
          <CardHeader className="py-2 px-3">
            <div className="flex justify-between items-center">
              <div className="flex items-center">
                <div className="h-8 w-8 rounded-full bg-teal-100 flex items-center justify-center mr-2">
                  <User className="h-4 w-4 text-teal-600" />
                </div>
                <div>
                  <CardTitle className="text-sm">
                    <DetailsDropdown
                      id={affiliate.affiliateId}
                      name={affiliate.name}
                    />
                  </CardTitle>
                  <CardDescription className="text-xs">
                    #{affiliate.affiliateId}
                  </CardDescription>
                </div>
              </div>
              <div className="text-right">
                <div className="text-sm font-medium">
                  {affiliate.commission}%
                </div>
                <div className="text-xs text-gray-500">{affiliate.country}</div>
              </div>
            </div>
          </CardHeader>
        </Card>
      </div>

      {expanded && hasChildren && (
        <div className="ml-7 pl-4 border-l border-gray-500 mt-2">
          {affiliate.children?.map((child) => (
            <TreeNode key={child.id} affiliate={child} level={level + 1} />
          ))}
        </div>
      )}
    </div>
  );
};

export function AffiliateTreeView({
  affiliateData,
}: {
  affiliateData: Affiliate;
}) {
  return <TreeNode affiliate={affiliateData} level={0} />;
}

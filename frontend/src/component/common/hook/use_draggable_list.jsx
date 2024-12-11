import { useState } from 'react';
import {
    DndContext,
    PointerSensor,
    useSensor,
    useSensors,
    closestCenter
} from '@dnd-kit/core';
import {
    SortableContext,
    useSortable,
    verticalListSortingStrategy,
    arrayMove
} from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';
import { List } from 'antd';
import { MenuOutlined } from '@ant-design/icons';

export default function useDraggableList(initialItems, onSortEnd) {
    const [items, setItems] = useState(initialItems);

    const sensors = useSensors(
        useSensor(PointerSensor, { activationConstraint: { distance: 5 } })
    );

    const handleDragEnd = (event) => {
        const { active, over } = event;
        if (!over || active.id === over.id) return;

        setItems((prevItems) => {
            const oldIndex = prevItems.findIndex((i) => i.id === active.id);
            const newIndex = prevItems.findIndex((i) => i.id === over.id);
            const newItems = arrayMove(prevItems, oldIndex, newIndex);

            if (onSortEnd) {
                onSortEnd(newItems.map((i) => i.id));
            }

            return newItems;
        });
    };

    function DraggableListProvider({ children }) {
        return (
            <DndContext
                sensors={sensors}
                collisionDetection={closestCenter}
                onDragEnd={handleDragEnd}
            >
                <SortableContext
                    items={items.map((i) => i.id)}
                    strategy={verticalListSortingStrategy}
                >
                    {/* We can wrap children in an AntD List for styling */}
                    <List
                        dataSource={[]}
                        renderItem={() => null}
                        style={{
                            background: '#fafafa',
                            padding: '10px',
                            borderRadius: '8px'
                        }}
                    >
                        {children}
                    </List>
                </SortableContext>
            </DndContext>
        );
    }

    function DraggableItem({ id, children }) {
        const { attributes, listeners, setNodeRef, transform, transition, isDragging } =
            useSortable({ id });
        const style = {
            transform: CSS.Transform.toString(transform),
            transition,
            background: isDragging ? '#f0f0f0' : '#fff',
            border: '1px solid #ddd',
            borderRadius: '6px',
            marginBottom: '8px'
        };

        return (
            <div ref={setNodeRef} style={style}>
                <List.Item
                    style={{ display: 'flex', alignItems: 'center', padding: 0 }}
                >
                    {/* Drag Handle on the left */}
                    <div
                        {...attributes}
                        {...listeners}
                        style={{
                            cursor: 'grab',
                            padding: '0 12px',
                            display: 'flex',
                            alignItems: 'center',
                            borderRight: '1px solid #ddd',
                        }}
                    >
                        <MenuOutlined />
                    </div>
                    <div style={{ flex: 1, padding: '8px 12px' }}>{children}</div>
                </List.Item>
            </div>
        );
    }

    return [items, setItems, DraggableListProvider, DraggableItem];
}

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
    horizontalListSortingStrategy,
    arrayMove
} from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';
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
                onSortEnd(newItems);
            }

            return newItems;
        });
    };

    function DraggableListProvider({ children, layout = 'vertical', fixedComponent }) {
        const isVertical = layout === 'vertical';
        return (
            <DndContext
                sensors={sensors}
                collisionDetection={closestCenter}
                onDragEnd={handleDragEnd}
            >
                <SortableContext
                    items={items.map((i) => i.id)}
                    strategy={
                        isVertical
                            ? verticalListSortingStrategy
                            : horizontalListSortingStrategy
                    }
                >
                    <div
                        style={{
                            display: isVertical ? 'block' : 'flex',
                            flexWrap: isVertical ? 'nowrap' : 'wrap',
                            // gap: '10px',
                            background: '#fafafa',
                            // padding: '10px',
                            borderRadius: '3px'
                        }}
                    >
                        {fixedComponent && (
                            <div
                                style={{
                                    margin: '10px',
                                    // padding: '10px',
                                    display: 'flex',
                                    alignItems: 'center'
                                }}
                            >
                                {fixedComponent}
                            </div>
                        )}
                        {children}
                    </div>
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
            borderRadius: '3px',
            display: 'flex',
            alignItems: 'center',
            cursor: 'grab',
            margin: '10px',
            height: '40px'
        };

        return (
            <div
                className="card"
                ref={setNodeRef}
                style={style}
                {...attributes}
                {...listeners}
            >
                <MenuOutlined style={{ marginRight: '10px' }} />
                <div>{children}</div>
            </div>
        );
    }

    return [items, setItems, DraggableListProvider, DraggableItem];
}

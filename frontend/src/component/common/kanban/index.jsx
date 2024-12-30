import React, { useState } from 'react';
import { DndContext, closestCenter, useDroppable, useDraggable } from '@dnd-kit/core';
import {
    arrayMove,
    SortableContext,
    verticalListSortingStrategy
} from '@dnd-kit/sortable';

const DraggableCard = ({ id, name }) => {
    const { attributes, listeners, setNodeRef, transform, transition } = useDraggable({
        id
    });

    const style = {
        transform: `translate(${transform?.x || 0}px, ${transform?.y || 0}px)`,
        transition,
        padding: '10px',
        margin: '5px 0',
        backgroundColor: 'white',
        border: '1px solid #ddd',
        borderRadius: '4px',
        cursor: 'grab'
    };

    return (
        <div ref={setNodeRef} style={style} {...listeners} {...attributes}>
            {name}
        </div>
    );
};

const DroppableColumn = ({ status, items, moveItem, onButtonClick }) => {
    const { setNodeRef } = useDroppable({
        id: status.value
    });

    return (
        <div
            ref={setNodeRef}
            style={{
                width: '300px',
                padding: '10px',
                marginRight: '10px', // Small gap between columns
                backgroundColor: '#f4f4f4',
                borderRadius: '4px',
                minHeight: '50px', // Minimum height for empty columns
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                boxSizing: 'border-box'
            }}
        >
            <div
                style={{
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    marginBottom: '10px'
                }}
            >
                <h3 style={{ textAlign: 'center', marginBottom: '5px' }}>
                    {status.label.toUpperCase()}
                </h3>
                <button
                    onClick={() => onButtonClick(status)}
                    style={{
                        padding: '5px 10px',
                        fontSize: '14px',
                        backgroundColor: '#007bff',
                        color: 'white',
                        border: 'none',
                        borderRadius: '4px',
                        cursor: 'pointer'
                    }}
                >
                    Add Item
                </button>
            </div>
            <SortableContext
                items={items.map((item) => item.id)}
                strategy={verticalListSortingStrategy}
            >
                {items.map((item) =>
                    item.status.id === status.value ? (
                        <DraggableCard key={item.id} id={item.id} name={item.title} />
                    ) : null
                )}
            </SortableContext>
        </div>
    );
};

const Kanban = ({ statusOption, data, onAdd }) => {
    const [items, setItems] = useState(data);

    const moveItem = (activeId, overId, newStatus) => {
        setItems((prevItems) => {
            const activeItem = prevItems.find((item) => item.id === activeId);
            const filteredItems = prevItems.filter((item) => item.id !== activeId);

            if (newStatus) {
                activeItem.status = newStatus;
            }

            const targetIndex = filteredItems.findIndex((item) => item.id === overId);
            const newItems =
                targetIndex >= 0
                    ? arrayMove(filteredItems, targetIndex, 0)
                    : filteredItems;

            return [activeItem, ...newItems];
        });
    };

    const handleDragEnd = (event) => {
        const { active, over } = event;

        if (active.id !== over?.id) {
            const activeItem = items.find((item) => item.id === active.id);
            const overItem = items.find((item) => item.id === over?.id);

            if (overItem?.status && activeItem?.status !== overItem.status) {
                moveItem(active.id, over.id, overItem.status);
            } else {
                moveItem(active.id, over.id);
            }
        }
    };
    return (
        <DndContext collisionDetection={closestCenter} onDragEnd={handleDragEnd}>
            <div
                style={{
                    display: 'flex',
                    overflowX: 'auto', // Enable horizontal scrolling
                    padding: '10px'
                }}
            >
                {statusOption.map((status) => (
                    <DroppableColumn
                        key={status.value}
                        status={status}
                        items={items}
                        moveItem={moveItem}
                        onButtonClick={onAdd}
                    />
                ))}
            </div>
        </DndContext>
    );
};

export default Kanban;
